package master

import (
	"github.com/benmizrahi/gobig/internal/common"
	"github.com/benmizrahi/gobig/internal/protos"
	dag "github.com/heimdalr/dag"
	_ "github.com/samber/lo"
)

type Mafream struct {
	dag     *dag.DAG
	root    string
	last    string
	context *Context
}

func NewDataFrame(c *Context, preplan *common.Maplan) *Mafream {

	dag := dag.NewDAG()
	root, _ := dag.AddVertex(preplan)

	return &Mafream{
		dag:     dag,
		root:    root,
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
	lastAction, _ := w.dag.AddVertex(&countPlan)
	w.dag.AddEdge(w.root, lastAction)

	_ = w.context.ExecuteDAG(w.root, lastAction, w.dag)

	// w.assignActions([]string{protos.COUNT})
	// results := w.context.DoAction(w.plan)
	// return lo.Sum(handleTasksResults[int](actions, results))

	return 0
}
