package bigfream

import "github.com/benmizrahi/gobig/internal/protos"

type Column struct {
	Type protos.DataType
	Name string
}

type BigOptions struct {
	Columns    []Column
	Repartiton int
}
