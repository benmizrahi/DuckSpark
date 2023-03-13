package worker

import (
	"net/http"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/benmizrahi/godist/protos"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Worker struct {
	MaxParallel int
	Master      string
	Host        string
	Port        int
	Http        *gin.Engine
}

func NewWorker(host string, port int) *Worker {
	return &Worker{
		MaxParallel: 10,
		Master:      "",
		Http:        gin.Default(),
		Host:        host,
		Port:        port,
	}
}

func (w *Worker) healthCheck(c *gin.Context) {
	res := &protos.HCRes{
		Uuid: uuid.New().String(),
		Time: timestamppb.Now(),
	}
	c.ProtoBuf(http.StatusOK, res)
}

func (w *Worker) Init() {
	w.Http.GET("/api/v1/health", w.healthCheck)
	w.Http.Run(":" + string(rune(w.Port)))
}

func (w *Worker) Stop() {
}
