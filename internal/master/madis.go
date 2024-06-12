package master

import (
	"github.com/benmizrahi/gobig/internal/common"
	"github.com/benmizrahi/gobig/internal/protos"
	"github.com/google/uuid"
	_ "github.com/samber/lo"
)

type Mafream struct {
	link    *common.LinkedList[common.Maplan]
	last    string
	context *Context
}

func NewDataFrame(c *Context, preplan *common.Maplan) *Mafream {
	return &Mafream{
		link:    common.NewLinkedList[common.Maplan](*preplan),
		context: c,
	}
}

func (w *Mafream) Show() *Mafream {
	return w
}

func (w *Mafream) Count() int {

	countPlan := common.Maplan{
		Action: protos.COUNT,
	}
	w.link.Push(countPlan)
	_ = w.context.ExecuteDAG(w.link, uuid.NewString())

	// w.assignActions([]string{protos.COUNT})
	// results := w.context.DoAction(w.plan)
	// return lo.Sum(handleTasksResults[int](actions, results))

	return 0
}
