package worker

import (
	"bytes"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/proto"

	"github.com/benmizrahi/gobig/internal/common"
	"github.com/benmizrahi/gobig/internal/protos"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Worker struct {
	ID          string
	MaxParallel int
	Master      string
	Host        string
	Port        int
	Plugins     map[string]common.IPluginContract
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
	go w.Http.Run(w.Host + ":" + strconv.Itoa(w.Port))
	logrus.Println("worker " + w.ID + " is listening at " + w.Host + ":" + strconv.Itoa(w.Port))
	return w
}

func (w *Worker) registerToMaster() {
	req := &protos.RegisterReq{
		Uuid: w.ID,
		Http: "http://" + w.Host + ":" + strconv.Itoa(w.Port),
	}
	body, err := proto.Marshal(req)
	if err != nil {
		logrus.Fatal(err)
	}
	_, err = http.Post(w.Master+"/api/register", "application/protobuf", bytes.NewReader(body))
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Println("Worker ID: " + w.ID + ",registered successfully to master")
}
