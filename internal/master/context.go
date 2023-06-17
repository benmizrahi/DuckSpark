package master

import (
	"bytes"
	"io"
	"net/http"
	"reflect"
	"sync"
	"time"

	"github.com/benmizrahi/godist/internal/plugins"
	"github.com/benmizrahi/godist/internal/plugins/contract"
	"github.com/benmizrahi/godist/internal/protos"
	"github.com/benmizrahi/godist/internal/worker"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

type Context struct {
	session *Master
	IsLocal bool
	Workers map[string]string
	Http    *gin.Engine
	plugins map[string]func() contract.IPluginContract
	plan    []*protos.IPartition
}

func NewContext(master *Master, isLocal bool, minWorkers int) *Context {

	context := &Context{
		IsLocal: isLocal,
		Workers: map[string]string{},
		session: master,
		Http:    gin.New(),
		plugins: map[string]func() contract.IPluginContract{},
	}

	context.Http.Use(ginlogrus.Logger(logrus.New()), gin.Recovery())
	context.Http.POST("/api/register", context.registerHandler)
	go context.Http.Run(master.MasterPath)
	log.Info("GoDist Master, master is listening on ", master.MasterPath)

	//load all internal plugins
	context.loadBuildInPlugins()

	//start all workers
	context.handleWorkers(minWorkers, isLocal, master.MasterPath)

	for len(context.Workers) != minWorkers {
		log.Info("GoDist Master, wating for %d workers to register..", minWorkers)
		time.Sleep(1 * time.Second)
	}

	log.Info("GoDist Master, all workers are ready")

	return context
}

func (c *Context) sendAyncTaskToWorker(worker string, partition *protos.IPartition) *protos.IPartitionResult {
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

func (c *Context) DoAction(plan []*protos.IPartition) []*protos.IPartitionResult {
	//TODO publish actions to workers
	var wg sync.WaitGroup

	allPartitionResults := []*protos.IPartitionResult{}
	keys := reflect.ValueOf(c.Workers).MapKeys()
	for index, partition := range plan {
		wg.Add(1)
		num := index % len(keys)
		worker := c.Workers[keys[num].String()]
		go func() {
			defer wg.Done()
			allPartitionResults = append(allPartitionResults, c.sendAyncTaskToWorker(worker, partition))
		}()
	}
	wg.Wait()

	return allPartitionResults
}

func (c *Context) loadBuildInPlugins() {
	for key, plugin := range plugins.MakeBuildIns() {
		c.plugins[key] = plugin
		log.Info("GoDist Master, plugin loaded ", key)
	}
}

func (c *Context) handleWorkers(minWorkers int, isLocal bool, masterPath string) {
	if isLocal {
		for i := 0; i < minWorkers; i++ {
			worker.NewWorker("localhost", 8080+i, masterPath)
		}
	} else {
		//TODO: implement GKE based orchstrations
	}
}
