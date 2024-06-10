package common

import "github.com/benmizrahi/gobig/internal/protos"

type Maplan struct {
	Action protos.Action
	Tasks  []*protos.Task
}
