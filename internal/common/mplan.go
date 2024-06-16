package common

import "github.com/benmizrahi/duckspark/internal/protos"

type Maplan struct {
	Action *string
	Plan   *string
	Tasks  []*protos.Task
}
