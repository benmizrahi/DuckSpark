package bigfream

import "github.com/benmizrahi/gobig/internal/domains"

type Column struct {
	Type domains.DataType
	Name string
}

type BigOptions struct {
	Columns    []Column
	Repartiton int
}
