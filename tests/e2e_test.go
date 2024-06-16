package tests

import (
	"testing"

	"github.com/benmizrahi/duckspark/internal/master"
)

func TestSQLAction(t *testing.T) {

	master.
		NewMaster(true, "localhost", 9999, 2).
		SQL("SELECT * FROM './resources/*.csv'")
}
