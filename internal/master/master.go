package master

import (
	"database/sql"
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
	rootPath string
	context  *Context
	http     *gin.Engine
	db       *sql.DB
}

func NewMaster(isLocal bool, host string, port int, minWorkers int) *Master {
	if masterInstance == nil {
		lock.Lock()
		m := &Master{
			rootPath: host + ":" + strconv.Itoa(port),
			http:     gin.New(),
		}

		m.http.Use(ginlogrus.Logger(logrus.New()), gin.Recovery())

		m.http.POST("/api/register", m.RegisterHandler)
		go m.http.Run(m.rootPath)
		log.Info("duckspark Master, master is listening on ", m.rootPath)

		m.context = NewContext(isLocal, minWorkers, m.rootPath)

		m.context.InitContext()

		connector, err := sql.Open("duckdb", "/tmp/m.db")
		if err != nil {
			logrus.Fatal("Unable to start worker err:", err)
		}
		m.db = connector

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

func (m *Master) SQL(query string) *Mafream {
	mapPlan := plugins.GetPlugin("duckdb_plugin").Plan(query, m.db)
	mf := NewDataFrame(m.context, &mapPlan)
	return mf.Show()
}
