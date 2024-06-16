package master

import (
	"io"
	"net/http"
	"strconv"
	"sync"

	"github.com/benmizrahi/duckspark/internal/plugins"
	"github.com/benmizrahi/duckspark/internal/protos"
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
	MasterPath string
	context    *Context
	Http       *gin.Engine
}

func NewMaster(isLocal bool, host string, port int, minWorkers int) *Master {
	if masterInstance == nil {
		lock.Lock()
		m := &Master{
			MasterPath: host + ":" + strconv.Itoa(port),
			Http:       gin.New(),
		}

		m.Http.Use(ginlogrus.Logger(logrus.New()), gin.Recovery())

		m.Http.POST("/api/register", m.RegisterHandler)
		go m.Http.Run(m.MasterPath)
		log.Info("duckspark Master, master is listening on ", m.MasterPath)

		m.context = NewContext(isLocal, minWorkers, m.MasterPath)

		m.context.InitContext()

		lock.Unlock()
		return m
	}
	return masterInstance
}

func (m *Master) RegisterHandler(c *gin.Context) {
	buf, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalln("Failed to parse register request:", err)
	}
	req := &protos.RegisterReq{}
	if err := proto.Unmarshal(buf, req); err != nil {
		log.Fatalln("Failed to parse register request:", err)
	}

	m.context.Workers[req.Uuid] = req.Http

	data := &protos.RegisterRes{
		Ok: true,
	}

	c.ProtoBuf(http.StatusOK, data)
}

func (m *Master) Load(path string) *Mafream {
	mapPlan := plugins.GetPlugin("fs_analyzer").Plan(path)
	mf := NewDataFrame(m.context, &mapPlan)
	return mf
}
