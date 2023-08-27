package tests

import (
	"testing"

	"github.com/benmizrahi/gobig/internal/bigfream"
	"github.com/benmizrahi/gobig/internal/protos"
)

func TestCountDynamicPartitons(t *testing.T) {
	wormup()

	data := [][]string{
		{"HELLO WORLD"},
		{"GOLANG IS"},
		{"THE BEST"},
		{"Programming Language"},
	}

	options := bigfream.BigOptions{
		Columns: []bigfream.Column{
			{
				Type: protos.DataType_string,
				Name: "Data",
			},
		},
	}

	count := gbigm.
		Parallelize(data, options).
		Count()

	if count != len(data) {
		t.Error()
	}
}
