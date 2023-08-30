package bigfream

import (
	"github.com/benmizrahi/gobig/internal/domains"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Bigfream struct {
	Columns    *[]Column
	context    *Context
	Options    BigOptions
	partitions []*domains.IPartition
}

func NewBigfream(c *Context, columns *[]Column) *Bigfream {
	return &Bigfream{
		Columns: columns,
		context: c,
	}
}

// func (w *Bigfream) Show(count int) {
// 	rowsPerPratition := int(count) / len(w.partitions)
// 	intWrapper := wrapperspb.Int32Value{Value: int32(rowsPerPratition)}
// 	intAny, _ := anypb.New(&intWrapper)
// 	w.assignActions([]string{protos.TAKE}, &[]*anypb.Any{intAny})
// 	results := w.context.DoAction(w.partitions)
// 	t := table.NewWriter()
// 	t.SetOutputMirror(os.Stdout)
// 	lo.ForEach(results, func(p *protos.IPartitionResult, index int) {
// 		lo.ForEach(p.TaskResults, func(tr *protos.TaskResult, index int) {
// 			lo.ForEach(tr.Rows, func(row *protos.Row, index int) {
// 				dataRow, _ := common.Deserialize[interface{}](row.CompressRow)
// 				t.AppendRow([]interface{}{dataRow})
// 			})
// 		})
// 	})
// 	t.Render()
// }

// func (w *Bigfream) Count() int {
// 	w.assignActions([]string{protos.COUNT}, &[]*anypb.Any{nil})
// 	results := w.context.DoAction(w.partitions)
// 	gtotal := lo.Map(results, func(p *protos.IPartitionResult, index int) int {
// 		partiton := lo.FlatMap(p.TaskResults, func(tr *protos.TaskResult, index int) []int {
// 			return lo.Map(tr.Rows, func(row *protos.Row, index int) int {
// 				dataRow, _ := common.Deserialize[int](row.CompressRow)
// 				return *dataRow
// 			})
// 		})
// 		return lo.Sum(partiton)
// 	})
// 	return lo.Sum(gtotal)
// }

func (w *Bigfream) Mapper(fun func(p domains.IPartition) *domains.IPartition) *Bigfream {
	// fmt.Println(signature(fun))
	return w
}

func (w *Bigfream) Reducer(func(p domains.IPartition)) *Bigfream {
	return w
}

func (w *Bigfream) AssginPartitons(partitions []*domains.IPartition) {
	w.partitions = partitions
}

func (w *Bigfream) assignActions(actions []string, params *[]*anypb.Any) {
	for _, partition := range w.partitions {
		partition.Tasks = append(partition.Tasks, &domains.Task{
			Uuid:              uuid.New().String(),
			Instruction:       actions,
			InstructionParams: *params,
			CreationTime:      timestamppb.Now(),
		})
	}
}
