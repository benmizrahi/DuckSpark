package master

import "github.com/gin-gonic/gin"

type Master struct {
	IsLocal bool
	Workers int
	Port    int
	Http    *gin.Engine
}

func NewMaster(isLocal bool) *Master {
	return &Master{
		IsLocal: isLocal,
		Workers: 2,
		Port:    9999,
		Http:    gin.Default(),
	}
}

func (w *Master) Init() {

}
