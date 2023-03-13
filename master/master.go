package master

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/benmizrahi/godist/protos"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
)

type Master struct {
	IsLocal bool
	Workers []map[string]string
	Port    int
	Host    string
	Http    *gin.Engine
}

func (w *Master) RegisterHandler(c *gin.Context) {
	buf, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalln("Failed to parse register request:", err)
	}
	req := &protos.RegisterReq{}
	if err := proto.Unmarshal(buf, req); err != nil {
		log.Fatalln("Failed to parse register request:", err)
	}

	c.ProtoBuf(http.StatusOK, protos.RegisterRes{
		Ok: true,
	})
}

func NewMaster(isLocal bool, host string, port int) *Master {

	return &Master{
		IsLocal: isLocal,
		Workers: []map[string]string{},
		Port:    port,
		Host:    host,
		Http:    gin.New(),
	}
}

func (w *Master) Init() {

	w.Http.POST("/api/register", w.RegisterHandler)
	w.Http.Run(w.Host + ":" + strconv.Itoa(w.Port))
}
