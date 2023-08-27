package master

import (
	"github.com/benmizrahi/gobig/internal/common"
	"github.com/benmizrahi/gobig/internal/protos"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
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

func (w *Mafream) Show(count int) {
	rowsPerPratition := int(count) / len(w.partitions)
	intWrapper := wrapperspb.Int32(int32(rowsPerPratition))
	intAny, _ := anypb.New(intWrapper)
	w.assignActions([]string{protos.TAKE}, &[]*anypb.Any{intAny})
	w.context.DoAction(w.partitions)
}

func (w *Mafream) Count() int {
	w.assignActions([]string{protos.COUNT}, nil)
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

func (w *Mafream) assignActions(actions []string, params *[]*anypb.Any) {
	for _, partition := range w.partitions {
		partition.Tasks = append(partition.Tasks, &protos.Task{
			Uuid:              uuid.New().String(),
			Instruction:       actions,
			InstructionParams: *params,
			CreationTime:      timestamppb.Now(),
		})
	}
}
