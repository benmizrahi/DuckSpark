package worker

import (
	"bytes"
	"log"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

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
	return &Worker{
		ID:          (uuid.New()).String(),
		MaxParallel: 10,
		Master:      "http://" + masterPath,
		Http:        gin.Default(),
		Host:        host,
		Port:        port,
	}
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

func (w *Worker) healthCheck(c *gin.Context) {
	res := &protos.HCRes{ 
		Uuid: uuid.New().String(),
		Time: timestamppb.Now(),
	}
	c.ProtoBuf(http.StatusOK, res)
}

func (w *Worker) Init() {

	w.registerToMaster()
	w.Http.GET("/api/v1/health", w.healthCheck)
	w.Http.Run(w.Host + ":" + strconv.Itoa(w.Port))
}
