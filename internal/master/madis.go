package master

import (
	"github.com/benmizrahi/gobig/internal/protos"
	"github.com/samber/lo"
	_ "github.com/samber/lo"
)

type Mafream struct {
	columns []string
	plan    []*protos.Task
	context *Context
}

func NewDataFrame(c *Context, columns []string, numPartitions int) *Mafream {
	partitions := make([]*protos.Task, numPartitions)
	for i := 0; i < numPartitions; i++ {
		partitions[i] = &protos.Task{}
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
	// results := w.context.DoAction(w.plan)
	// w.handleTasksResults(actions, results)
	return w
}

func (w *Mafream) Count() int {
	actions := []string{protos.COUNT}
	w.assignActions([]string{protos.COUNT})
	results := w.context.DoAction(w.plan)
	return lo.Sum(handleTasksResults[int](actions, results))
}

func (w *Mafream) assignActions(actions []string) {

	// for _, partition := range w.plan {
	// 	partition.Tasks = append(partition.Tasks, &protos.Task{
	// 		Uuid:         uuid.New().String(),
	// 		Instactions:  actions,
	// 		CreationTime: timestamppb.Now(),
	// 	})
	// }
}

func handleTasksResults[T any](actions []string, results []*protos.TaskResult) []T {
	// for _, res := range results {
	// 	res.TaskResults
	// }

	return nil
}
