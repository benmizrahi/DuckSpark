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
	action := "COUNT"
	countPlan := common.Maplan{
		Action: &action,
		Tasks: []*protos.Task{{
			Uuid:         uuid.New().String(),
			Commands:     []string{" COUNT"},
			CreationTime: timestamppb.Now()},
		},
	}

	w.link.Push(countPlan)
	_ = w.context.ExecuteAction(uuid.NewString(), w.link)

	return 0
}
