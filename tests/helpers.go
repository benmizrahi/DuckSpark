package tests

import "github.com/benmizrahi/gobig/internal/master"

var gbigm *master.Master

func wormup() {
	if gbigm == nil {
		gbigm = master.
			NewMaster(true, "localhost", 9999, 2)
	}
}
