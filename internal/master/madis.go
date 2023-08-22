package master

import (
	"github.com/benmizrahi/gobig/internal/protos"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Mafream struct {
	columns []string
	plan    []*protos.IPartition
	context *Context
}

func NewDataFrame(c *Context, columns []string, numPartitions int) *Mafream {
	partitions := make([]*protos.IPartition, numPartitions)
	for i := 0; i < numPartitions; i++ {
		partitions[i] = &protos.IPartition{}
	}

	return &Mafream{
		columns: columns,
		plan:    partitions,
		context: c,
	}
}

func (w *Mafream) Show() *Mafream {
	actions := []string{protos.TAKE, protos.LIMIT}
	w.assignActions(actions)
	results := w.context.DoAction(w.plan)
	w.handleTasksResults(actions, results)
	return w
}

func (w *Mafream) Count() *Mafream {
	actions := []string{protos.COUNT}
	w.assignActions([]string{protos.COUNT})
	results := w.context.DoAction(w.plan)
	w.handleTasksResults(actions, results)
	return w
}

func (w *Mafream) assignActions(actions []string) {
	for _, partition := range w.plan {
		partition.Tasks = append(partition.Tasks, &protos.Task{
			Uuid:         uuid.New().String(),
			Instactions:  actions,
			CreationTime: timestamppb.Now(),
		})
	}
}

func (w *Mafream) handleTasksResults(actions []string, res []*protos.IPartitionResult) {

}
