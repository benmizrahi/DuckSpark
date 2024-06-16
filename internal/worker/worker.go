package worker

import (
	"bytes"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/proto"

	"github.com/benmizrahi/duckspark/internal/protos"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/marcboeker/go-duckdb"
	"github.com/sirupsen/logrus"
)

type Worker struct {
	ID     string
	Master string
	Host   string
	Port   int
	Http   *gin.Engine
	db     *sql.DB
}

func NewWorker(host string, port int, masterPath string) *Worker {

	w := &Worker{
		ID:     (uuid.New()).String(),
		Master: "http://" + masterPath,
		Http:   gin.Default(),
		Host:   host,
		Port:   port,
	}

	connector, err := sql.Open("duckdb", "/tmp/w-"+w.ID+".db")
	if err != nil {
		logrus.Fatal("Unable to start worker err:", err)
	}
	w.db = connector

	w.registerToMaster()

	w.Http.GET("/api/v1/health", w.healthCheck)
	w.Http.POST("/api/v1/tasks", w.taskHandler)

	go w.Http.Run(w.Host + ":" + strconv.Itoa(w.Port))

	logrus.Info("worker %s is listening at %s", w.ID, w.Host+":"+strconv.Itoa(w.Port))
	return w
}

func (w *Worker) registerToMaster() {
	req := &protos.RegisterReq{
		Uuid: w.ID,
		Http: "http://" + w.Host + ":" + strconv.Itoa(w.Port),
	}
	body, err := proto.Marshal(req)
	_, err = http.Post(w.Master+"/api/register", "application/protobuf", bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
}
