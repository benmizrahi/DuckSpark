package master

import (
	"github.com/benmizrahi/godist/internal/protos"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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
	for _, partition := range w.plan {
		partition.Tasks = append(partition.Tasks, &protos.Task{
			Uuid:         uuid.New().String(),
			Instactions:  []string{protos.TAKE, protos.LIMIT},
			CreationTime: timestamppb.Now(),
		})
	}

	planResults := w.context.DoAction(w.plan)
	for _, res := range planResults {
		logrus.Info(res.TaskResults)
	}
	return w
}

func (w *Mafream) Count() *Mafream {
	return w
}
