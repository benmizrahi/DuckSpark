package tests

import (
	"testing"

	"github.com/benmizrahi/gobig/internal/common"
)

func TestCountDynamicPartitons(t *testing.T) {
	wormup()

	data := [][]string{
		{"HELLO WORLD"},
		{"GOLANG IS"},
		{"THE BEST"},
		{"Programming Language"},
	}

	options := common.Options{
		Columns: []string{
			"data",
		},
	}

	count := gbigm.
		Parallelize(data, options).
		Count()
	if count != len(data) {
		t.Error()
	}
}

func TestCountStaticPartitonsLargerThenCalculated(t *testing.T) {
	wormup()

	data := [][]string{
		{"HELLO WORLD"},
		{"GOLANG IS"},
		{"THE BEST"},
		{"Programming Language"},
	}

	options := common.Options{
		Repartiton: 4,
		Columns: []string{
			"data",
		},
	}

	count := gbigm.
		Parallelize(data, options).
		Count()

	if count != len(data) {
		t.Error()
	}

}

func TestCountStaticPartitonsSmallerThenCalculated(t *testing.T) {
	wormup()

	data := [][]string{
		{"HELLO WORLD"},
		{"GOLANG IS"},
		{"THE BEST"},
		{"Programming Language"},
	}

	options := common.Options{
		Repartiton: 1,
		Columns: []string{
			"data",
		},
	}

	count := gbigm.
		Parallelize(data, options).
		Count()

	if count != len(data) {
		t.Error()
	}
}
