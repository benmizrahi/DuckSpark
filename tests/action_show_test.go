package tests

import (
	"testing"

	"github.com/benmizrahi/gobig/internal/common"
)

func TestShowFunction(t *testing.T) {
	wormup()

	data := [][]string{
		{"HELLO WORLD"},
		{"GOLANG IS"},
		{"THE BEST"},
		{"Programming Language"},
	}

	options := common.Options{
		Repartiton: 10,
		Columns: []string{
			"data",
		},
	}

	gbigm.
		Parallelize(data, options).
		Show(10)

}
