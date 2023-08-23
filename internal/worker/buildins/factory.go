package buildins

import (
	"github.com/benmizrahi/gobig/internal/protos"
)

func MakeTaskInstruction(partition *protos.IPartition, t *protos.Task) []*protos.TaskResult {
	res := []*protos.TaskResult{}
	for _, instruction := range t.Instruction {
		switch instruction {
		case protos.IN_MEMORY_READ:
			continue
		case protos.COUNT:
			continue
		case protos.LIMIT:
			continue
		case protos.TAKE:
			continue
		}
	}
	return res
}
