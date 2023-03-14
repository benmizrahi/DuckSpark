package master

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/benmizrahi/godist/plugins"
	"github.com/benmizrahi/godist/protos"
	"github.com/benmizrahi/godist/worker"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
)

type Master struct {
	IsLocal    bool
	Workers    map[string]string
	Plugins    map[string]plugins.IPluginContract
	MasterPath string
	context    *Context
	Http       *gin.Engine
}

func handleWorkers(minWorkers int, isLocal bool, masterPath string) {
	if isLocal {
		for i := 0; i < minWorkers; i++ {
			worker.NewWorker("localhost", 999+i, masterPath)
		}
	} else {
		//TODO: implement GKE based orchstrations
	}
}

func NewMaster(isLocal bool, host string, port int, minWorkers int) *Master {
	w := &Master{
		IsLocal:    isLocal,
		Workers:    map[string]string{},
		Plugins:    map[string]plugins.IPluginContract{},
		MasterPath: host + ":" + strconv.Itoa(port),
		Http:       gin.New(),
	}

	w.Http.POST("/api/register", w.registerHandler)
	go w.Http.Run(w.MasterPath)

	handleWorkers(minWorkers, isLocal, w.MasterPath)

	for len(w.Workers) != minWorkers {
		time.Sleep(1 * time.Second)
	}

	return w
}

func (w *Master) registerHandler(c *gin.Context) {
	buf, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalln("Failed to parse register request:", err)
	}
	req := &protos.RegisterReq{}
	if err := proto.Unmarshal(buf, req); err != nil {
		log.Fatalln("Failed to parse register request:", err)
	}

	w.Workers[req.Uuid] = req.Uuid

	data := &protos.RegisterRes{
		Ok: true,
	}

	c.ProtoBuf(http.StatusOK, data)
}

func (w *Master) LoadPlugin(plugin plugins.IPluginContract) {
	w.Plugins[plugin.Name()] = plugin
}

func (w *Master) Context() *Context {
	if w.context == nil {
		//TODO make it thread-safe
		w.context = NewContext(w)
	}
	return w.context
}
