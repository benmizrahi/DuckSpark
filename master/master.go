package master

import (
	"bytes"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"sync"
	"time"

	"github.com/benmizrahi/godist/plugins"
	"github.com/benmizrahi/godist/plugins/contract"
	"github.com/benmizrahi/godist/protos"
	"github.com/benmizrahi/godist/worker"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

// singeltone instance of master!
var lock = &sync.Mutex{}

// Singel instance
var masterInstance *Master

type Master struct {
	IsLocal    bool
	Workers    map[string]string
	Plugins    map[string]func() contract.IPluginContract
	MasterPath string
	context    *Context
	Http       *gin.Engine
}

func NewMaster(isLocal bool, host string, port int, minWorkers int) *Master {
	if masterInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		log.Info("GoDist Master, Creating new master instance")
		w := &Master{
			IsLocal:    isLocal,
			Workers:    map[string]string{},
			Plugins:    map[string]func() contract.IPluginContract{},
			MasterPath: host + ":" + strconv.Itoa(port),
			Http:       gin.New(),
		}

		w.Http.Use(ginlogrus.Logger(logrus.New()), gin.Recovery())
		w.Http.POST("/api/register", w.registerHandler)
		go w.Http.Run(w.MasterPath)
		log.Info("GoDist Master, master is listening on ", w.MasterPath)

		//load all internal plugins
		w.loadBuildInPlugins()

		//start all workers
		w.handleWorkers(minWorkers, isLocal, w.MasterPath)

		for len(w.Workers) != minWorkers {
			log.Info("GoDist Master, wating for %d workers to register..", minWorkers)
			time.Sleep(1 * time.Second)
		}

		log.Info("GoDist Master, all workers are ready")
		return w
	}
	return masterInstance
}

func (w *Master) handleWorkers(minWorkers int, isLocal bool, masterPath string) {
	if isLocal {
		for i := 0; i < minWorkers; i++ {
			worker.NewWorker("localhost", 999+i, masterPath)
		}
	} else {
		//TODO: implement GKE based orchstrations
	}
}

func (w *Master) loadBuildInPlugins() {
	for key, plugin := range plugins.MakeBuildIns() {
		w.Plugins[key] = plugin
		log.Info("GoDist Master, plugin loaded ", key)
	}
}

func (w *Master) sendAyncTaskToWorker(worker string, partition *protos.IPartition) *protos.IPartitionResult {
	body, err := proto.Marshal(partition)
	if err != nil {
		log.Fatal("error:", err)
	}
	res, err := http.Post(worker+"/api/v1/tasks", "application/protobuf", bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	buf, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	result := protos.IPartitionResult{}
	err = proto.Unmarshal(buf, &result)
	if err != nil {
		log.Fatal(err)
	}

	return &result
}

func (w *Master) DoAction(plan []protos.IPartition) bool {
	//TODO publish actions to workers
	var wg sync.WaitGroup
	keys := reflect.ValueOf(w.Workers).MapKeys()
	for index, partition := range plan {
		wg.Add(1)
		num := index % len(keys)
		worker := w.Workers[keys[num].String()]
		go func(master *Master, w string, p protos.IPartition, wg sync.WaitGroup) {
			master.sendAyncTaskToWorker(w, &p)
			defer wg.Done()
		}(w, worker, partition, wg)
	}
	wg.Wait()

	return false
}

func (w *Master) Context() *Context {
	if w.context == nil {
		lock.Lock()
		defer lock.Unlock()
		w.context = NewContext(w)
	}
	return w.context
}
