package tests

// import (
// 	"testing"

// 	"github.com/benmizrahi/gobig/internal/bigfream"
// 	"github.com/benmizrahi/gobig/internal/protos"
// )

// func TestCountDynamicPartitons(t *testing.T) {
// 	wormup()

// 	data := [][]string{
// 		{"HELLO WORLD"},
// 		{"GOLANG IS"},
// 		{"THE BEST"},
// 		{"Programming Language"},
// 	}

// 	options := bigfream.BigOptions{
// 		Columns: []bigfream.Column{
// 			{
// 				Type: protos.DataType_string,
// 				Name: "Data",
// 			},
// 		},
// 	}

// 	gbigm.
// 		Parallelize(data, options).
// 		Mapper(func(p protos.IPartition) *protos.IPartition {
// 			return nil
// 		}).
// 		Reducer(func(p protos.IPartition) {

// 		})
// }
