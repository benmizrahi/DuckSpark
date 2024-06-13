package worker

import (
	"io"
	"log"
	"net/http"

	"github.com/benmizrahi/duckspark/internal/plugins"
	"github.com/benmizrahi/duckspark/internal/protos"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (w *Worker) healthCheck(c *gin.Context) {
	res := &protos.HCRes{
		Uuid: uuid.New().String(),
		Time: timestamppb.Now(),
	}
	c.ProtoBuf(http.StatusOK, res)
}

func (w *Worker) taskHandler(c *gin.Context) {
	buf, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalln("Failed to parse register request:", err)
	}

	task := &protos.Task{}
	if err := proto.Unmarshal(buf, task); err != nil {
		log.Fatalln("Failed to parse register request:", err)
	}

	tResult := plugins.GetPlugin(task.Plugin).Execute(task)

	if !tResult.Dataflow {
		CacheIt(task.DagId, tResult.Data)
		tResult.Data = nil
	}

	c.ProtoBuf(http.StatusOK, tResult)
}
