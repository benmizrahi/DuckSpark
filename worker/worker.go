package worker

import (
	"net/http"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/benmizrahi/godistodist/proto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Worker struct {
	MaxParallel int
	Master      string
	Port        int
	Http        *gin.Engine
}

func NewWorker(master string) *Worker {
	return &Worker{
		Port:        9999,
		MaxParallel: 10,
		Master:      `http://` + master + `:9991/api`,
		Http:        gin.Default(),
	}
}

func (w *Worker) Init() {
	w.Http.GET("/api/v1/health", func(c *gin.Context) {
		res := &proto.HCRes{
			Uuid: uuid.New().String(),
			Time: timestamppb.Now(),
		}
		c.ProtoBuf(http.StatusOK, res)
	})
}

func (w *Worker) Start() {
	w.Http.Run(":" + string(rune(w.Port)))
}

func (w *Worker) Stop() {
}
