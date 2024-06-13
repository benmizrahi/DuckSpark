package master

import (
	"bytes"
	"io"
	"net/http"
	"reflect"
	"sync"
	"time"

	"github.com/benmizrahi/duckspark/internal/common"
	"github.com/benmizrahi/duckspark/internal/protos"
	"github.com/benmizrahi/duckspark/internal/worker"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

type Context struct {
	IsLocal bool
	Workers map[string]string
	//private
	minWorkers int
	masterPath string
}

func NewContext(isLocal bool, minWorkers int, masterPath string) *Context {

	context := &Context{
		IsLocal:    isLocal,
		Workers:    map[string]string{},
		minWorkers: minWorkers,
		masterPath: masterPath,
	}
	return context
}

func (c *Context) InitContext() {
	// start all workers
	c.handleWorkers(c.minWorkers, c.IsLocal, c.masterPath)

	for len(c.Workers) != c.minWorkers {
		log.Info("duckspark Master, wating for %d workers to register..", c.minWorkers)
		time.Sleep(1 * time.Second)
	}
	log.Info("duckspark Master, all workers are ready")

}

func (c *Context) sendAyncTaskToWorker(worker string, partition *protos.Task) *protos.TaskResult {
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
	result := protos.TaskResult{}
	err = proto.Unmarshal(buf, &result)
	if err != nil {
		log.Fatal(err)
	}

	return &result
}

func (c *Context) ExecuteAction(actionId string, list *common.LinkedList[common.Maplan]) []*protos.TaskResult {

	tasksResults := []*protos.TaskResult{}
	maplan, exits := list.Pop()
	for {
		if !exits {
			break
		}
		tasksResults = c.distributeTasksToWorkers(actionId, maplan)
		maplan, exits = list.Pop()
	}

	return tasksResults
}

func (c *Context) distributeTasksToWorkers(actionId string, maplan common.Maplan) []*protos.TaskResult {
	allTasksResults := []*protos.TaskResult{}
	var wg sync.WaitGroup

	keys := reflect.ValueOf(c.Workers).MapKeys()
	for index, task := range maplan.Tasks {
		task.DagId = actionId
		wg.Add(1)
		num := index % len(keys)
		worker := c.Workers[keys[num].String()]
		go func() {
			defer wg.Done()
			allTasksResults = append(allTasksResults, c.sendAyncTaskToWorker(worker, task))
		}()
	}

	return allTasksResults
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
