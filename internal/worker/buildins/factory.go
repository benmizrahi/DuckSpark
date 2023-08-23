package buildins

import (
	"github.com/benmizrahi/gobig/internal/protos"
)

func MakeTaskInstruction(partition *protos.IPartition, t *protos.Task) []*protos.TaskResult {
	res := []*protos.TaskResult{}
	for _, instruction := range t.Instruction {
		switch instruction {
		case protos.COUNT:
			res = append(res, Count(t.Uuid, partition.Data))
		}
	}
	return res
}
