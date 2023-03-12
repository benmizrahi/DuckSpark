package worker

import (
	"github.com/gin-gonic/gin"
)

type Worker struct {
	MaxParallel int
	Master      string
	Port        int
}

func NewWorker() *Worker {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {})
	return &Worker{}
}
