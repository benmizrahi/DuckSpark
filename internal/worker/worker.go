package worker

import (
	"bytes"
	"net/http"
	"strconv"

	"github.com/benmizrahi/gobig/internal/common"
	"github.com/benmizrahi/gobig/internal/domains"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

type Worker struct {
	ID          string
	MaxParallel int
	Master      string
	Host        string
	Port        int
	Plugins     map[string]common.IPluginContract
}

func NewWorker(port int, masterPath string) *Worker {

	w := &Worker{
		ID:          (uuid.New()).String(),
		MaxParallel: 10,
		Master:      "http://" + masterPath,
		Port:        port,
	}

	w.registerToMaster()

	logrus.Println("worker " + w.ID + " is listening at " + w.Host + ":" + strconv.Itoa(w.Port))

	return w
}

func (w *Worker) registerToMaster() {
	req := &domains.RegisterReq{
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
