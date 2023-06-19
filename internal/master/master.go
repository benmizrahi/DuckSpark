package master

import (
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
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
		log.Info("GoDist Master, Creating new master instance")
		m := &Master{
			MasterPath: host + ":" + strconv.Itoa(port),
			Http:       gin.New(),
		}

		m.Http.Use(ginlogrus.Logger(logrus.New()), gin.Recovery())
		m.Http.POST("/api/register", m.context.RegisterHandler)
		go m.Http.Run(m.MasterPath)
		log.Info("GoDist Master, master is listening on ", m.MasterPath)

		m.context = NewContext(isLocal, minWorkers, m.MasterPath)
		lock.Unlock()
		return m
	}
	return masterInstance
}

func (m *Master) Parallelize(csv string) *Mafream {
	mf := NewDataFrame(m.context, []string{}, 1)
	return mf
}

func (m *Master) Load() *Mafream {

	mf := NewDataFrame(m.context, []string{}, 1)
	return mf
}
