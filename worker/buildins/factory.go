package buildins

import (
	"github.com/benmizrahi/godist/protos"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MakeInstactions(t *protos.Task) *protos.TaskResult {
	res := protos.TaskResult{
		Uuid:    t.Uuid,
		Status:  true,
		EndTime: timestamppb.Now(),
	}
	return &res
}
