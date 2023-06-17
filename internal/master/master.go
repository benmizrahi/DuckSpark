package master

import (
	"strconv"
	"sync"

	"github.com/benmizrahi/godist/internal/protos"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// singeltone instance of master!
var lock = &sync.Mutex{}

// Singel instance
var masterInstance *Master

type Master struct {
	MasterPath string
	context    *Context
}

func NewMaster(isLocal bool, host string, port int, minWorkers int) *Master {
	if masterInstance == nil {
		lock.Lock()
		log.Info("GoDist Master, Creating new master instance")
		w := &Master{
			MasterPath: host + ":" + strconv.Itoa(port),
		}
		w.context = NewContext(w, isLocal, minWorkers)
		lock.Unlock()
		return w
	}
	return masterInstance
}

func (w *Master) Extract(from string, config map[string]string) *Master {
	return w
}

func (w *Master) Parallelize() *Master {
	return w
}

func (w *Master) Transform() *Master {
	return w
}

func (w *Master) Load(job string) *Master {
	return w
}

func (w *Master) Show() *Master {
	for _, partition := range w.context.plan {
		partition.Tasks = append(partition.Tasks, &protos.Task{
			Uuid:         uuid.New().String(),
			Instactions:  []string{protos.TAKE, protos.LIMIT},
			CreationTime: timestamppb.Now(),
		})
	}

	planResults := w.context.DoAction(w.context.plan)
	for _, res := range planResults {
		logrus.Info(res.TaskResults)
	}
	return w
}
