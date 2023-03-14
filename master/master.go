package master

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/benmizrahi/godist/plugins"
	"github.com/benmizrahi/godist/protos"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
)

type Master struct {
	IsLocal bool
	Workers []map[string]string
	Plugins map[string]plugins.IPluginContract
	Port    int
	Host    string
	Http    *gin.Engine
}

func NewMaster(isLocal bool, host string, port int) *Master {

	return &Master{
		IsLocal: isLocal,
		Workers: []map[string]string{},
		Plugins: map[string]plugins.IPluginContract{},
		Port:    port,
		Host:    host,
		Http:    gin.New(),
	}
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

	data := &protos.RegisterRes{
		Ok: true,
	}

	c.ProtoBuf(http.StatusOK, data)
}

func (w *Master) Start() *Master {

	w.Http.POST("/api/register", w.registerHandler)
	w.Http.Run(w.Host + ":" + strconv.Itoa(w.Port))

	return w
}

func (w *Master) LoadPlugin(plugin plugins.IPluginContract) {
	w.Plugins[plugin.Name()] = plugin
}

func (w *Master) Extract(job string) *Master {
	return w
}

func (w *Master) Transform(job string) *Master {
	return w
}

func (w *Master) Load(job string) *Master {
	return w
}
