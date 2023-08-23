package master

import (
	"github.com/benmizrahi/gobig/internal/protos"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Mafream struct {
	columns    []string
	partitions []*protos.IPartition
	context    *Context
}

func NewDataFrame(c *Context, columns []string) *Mafream {
	return &Mafream{
		columns: columns,
		context: c,
	}
}

func (w *Mafream) Show() *Mafream {
	actions := []string{protos.TAKE, protos.LIMIT}
	w.assignActions(actions)
	results := w.context.DoAction(w.partitions)
	w.handleTasksResults(actions, results)
	return w
}

func (w *Mafream) Count() *Mafream {
	actions := []string{protos.COUNT}
	w.assignActions([]string{protos.COUNT})
	results := w.context.DoAction(w.partitions)
	w.handleTasksResults(actions, results)
	return w
}

func (w *Mafream) assignActions(actions []string) {
	for _, partition := range w.partitions {
		partition.Tasks = append(partition.Tasks, &protos.Task{
			Uuid:         uuid.New().String(),
			Instactions:  actions,
			CreationTime: timestamppb.Now(),
		})
	}
}

func (w *Mafream) handleTasksResults(actions []string, res []*protos.IPartitionResult) {

}
