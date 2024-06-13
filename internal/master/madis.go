package master

import (
	"github.com/benmizrahi/duckspark/internal/common"
	"github.com/benmizrahi/duckspark/internal/protos"
	"github.com/google/uuid"
	_ "github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Mafream struct {
	link    *common.LinkedList[common.Maplan]
	last    string
	context *Context
}

func NewDataFrame(c *Context, preplan *common.Maplan) *Mafream {
	return &Mafream{
		link:    common.NewLinkedList(*preplan),
		context: c,
	}
}

func (w *Mafream) Show() *Mafream {
	return w
}

func (w *Mafream) Count() int {

	logrus.Info("trigger action Count")

	count_task := &protos.Task{
		Uuid:         uuid.New().String(),
		Instactions:  []string{"COUNT"},
		CreationTime: timestamppb.Now(),
	}

	countPlan := common.Maplan{
		Action: common.COUNT,
		Tasks: []*protos.Task{
			count_task,
		},
	}

	w.link.Push(countPlan)
	_ = w.context.ExecuteAction(uuid.NewString(), w.link)

	// w.assignActions([]string{protos.COUNT})
	// results := w.context.DoAction(w.plan)
	// return lo.Sum(handleTasksResults[int](actions, results))

	return 0
}
