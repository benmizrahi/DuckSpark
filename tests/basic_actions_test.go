package tests

import (
	"testing"

	"github.com/benmizrahi/gobig/internal/master"
)

func TestCountCSV(t *testing.T) {
	results := master.
		NewMaster(true, "localhost", 9999, 2).
		Load("./resources/").
		Count()

	if results != 3 {
		t.Fatal("count fucntion produce wrong awnswer")
	}
}
