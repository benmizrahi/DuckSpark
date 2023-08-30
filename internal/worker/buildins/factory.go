package buildins

import (
	"github.com/benmizrahi/gobig/internal/domains"
)

func MakeTaskInstruction(partition *domains.IPartition, t *domains.Task) []*domains.TaskResult {
	res := []*domains.TaskResult{}
	for _, instruction := range t.Instruction {
		switch instruction {
		case domains.DataType_string.String():
		}
	}
	return res
}
