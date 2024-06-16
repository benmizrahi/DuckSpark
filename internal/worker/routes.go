package worker

import (
	"context"
	"io"
	"log"
	"net/http"

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

	for _, commend := range task.Commands {

		rows, err := w.db.QueryContext(context.Background(), commend)
		if err != nil {
			log.Fatalln("Failed to execute commend", err)
		}
		rows.Close()
		CacheIt(task.StageId, &rows)
	}

	c.ProtoBuf(http.StatusOK, &protos.TaskResult{
		Uuid:    task.StageId,
		Status:  true,
		EndTime: timestamppb.Now(),
		Data:    []*protos.DataRow{},
	})
}
