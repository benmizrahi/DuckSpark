package master

import (
	"github.com/benmizrahi/gobig/internal/common"
	"github.com/benmizrahi/gobig/internal/protos"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
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
	return w
}

func (w *Mafream) Collect() []*protos.Row {
	//TODO: implement map
	return []*protos.Row{}
}

func (w *Mafream) Count() int {
	w.assignActions([]string{protos.COUNT})
	results := w.context.DoAction(w.partitions)
	gtotal := 0
	for _, p := range results {
		for _, t := range p.TaskResults {
			data := lo.Map(t.Rows, func(d *protos.Row, index int) int {
				count, err := common.Deserialize(d.CompressRow)
				if err != nil {
					logrus.Error("error deserialize data,", err)
				}
				total := lo.Reduce(count, func(agg int, item interface{}, _ int) int {
					return agg + item.(int)
				}, 0)
				return total
			})
			gtotal += lo.Reduce(data, func(agg int, item int, _ int) int {
				return agg + item
			}, 0)
		}
	}
	return gtotal
}

func (w *Mafream) assignActions(actions []string) {
	for _, partition := range w.partitions {
		partition.Tasks = append(partition.Tasks, &protos.Task{
			Uuid:         uuid.New().String(),
			Instruction:  actions,
			CreationTime: timestamppb.Now(),
		})
	}
}
