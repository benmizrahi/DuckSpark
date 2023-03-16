package worker

import (
	"bytes"
	"log"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/proto"

	"github.com/benmizrahi/godist/protos"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Worker struct {
	ID          string
	MaxParallel int
	Master      string
	Host        string
	Port        int
	Http        *gin.Engine
}

func NewWorker(host string, port int, masterPath string) *Worker {

	w := &Worker{
		ID:          (uuid.New()).String(),
		MaxParallel: 10,
		Master:      "http://" + masterPath,
		Http:        gin.Default(),
		Host:        host,
		Port:        port,
	}

	w.registerToMaster()
	w.Http.GET("/api/v1/health", w.healthCheck)
	w.Http.POST("/api/v1/tasks", w.tasksHandler)
	w.Http.Run(w.Host + ":" + strconv.Itoa(w.Port))

	return w
}

func (w *Worker) registerToMaster() {
	req := &protos.RegisterReq{
		Uuid: w.ID,
	}
	body, err := proto.Marshal(req)
	_, err = http.Post(w.Master+"/api/register", "application/protobuf", bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
}
