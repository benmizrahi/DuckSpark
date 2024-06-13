package common

import "github.com/benmizrahi/duckspark/internal/protos"

type Maplan struct {
	Action Action
	Plan   Planner
	Tasks  []*protos.Task
}
