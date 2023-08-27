package buildins

import (
	"github.com/benmizrahi/gobig/internal/protos"
	"google.golang.org/protobuf/types/known/anypb"
)

func Take(uuid string, data []*protos.Row, params []*anypb.Any) *protos.TaskResult {
	res := protos.TaskResult{
		Uuid:   uuid,
		Status: true,
		Rows:   []*protos.Row{},
	}

	return &res
}
