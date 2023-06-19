package master

import (
	"bytes"
	"io"
	"net/http"
	"reflect"
	"sync"
	"time"

	"github.com/benmizrahi/godist/internal/protos"
	"github.com/benmizrahi/godist/internal/worker"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

type Context struct {
	IsLocal bool
	Workers map[string]string
}

func NewContext(isLocal bool, minWorkers int, masterPath string) *Context {

	context := &Context{
		IsLocal: isLocal,
		Workers: map[string]string{},
	}

	//start all workers
	context.handleWorkers(minWorkers, isLocal, masterPath)

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

func (c *Context) handleWorkers(minWorkers int, isLocal bool, masterPath string) {
	if isLocal {
		for i := 0; i < minWorkers; i++ {
			worker.NewWorker("localhost", 8080+i, masterPath)
		}
	} else {
		//TODO: implement GKE based orchstrations
	}
}

func (co *Context) RegisterHandler(c *gin.Context) {
	buf, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalln("Failed to parse register request:", err)
	}
	req := &protos.RegisterReq{}
	if err := proto.Unmarshal(buf, req); err != nil {
		log.Fatalln("Failed to parse register request:", err)
	}

	co.Workers[req.Uuid] = req.Http

	data := &protos.RegisterRes{
		Ok: true,
	}

	c.ProtoBuf(http.StatusOK, data)
}
