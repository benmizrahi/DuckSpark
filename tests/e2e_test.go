package tests

import (
	"testing"

	"github.com/benmizrahi/duckspark/internal/master"
)

func TestCountAction(t *testing.T) {

	master.
		NewMaster(true, "localhost", 9999, 2).
		Load("./resources/").
		Count()
}
