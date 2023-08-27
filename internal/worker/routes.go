package worker

import (
	"io"
	"log"
	"net/http"

	"github.com/benmizrahi/gobig/internal/protos"
	"github.com/benmizrahi/gobig/internal/worker/buildins"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (w *Worker) healthCheck(c *gin.Context) {
	res := &protos.HCRes{
		Uuid: uuid.New().String(),
		Time: timestamppb.Now(),
	}
	c.ProtoBuf(http.StatusOK, res)
}

func (w *Worker) tasksHandler(c *gin.Context) {
	buf, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalln("Failed to parse register request:", err)
	}

	partition := &protos.IPartition{}
	if err := proto.Unmarshal(buf, partition); err != nil {
		log.Fatalln("Failed to parse register request:", err)
	}

	res := &protos.IPartitionResult{
		TaskResults: []*protos.TaskResult{},
	}

	for _, task := range partition.Tasks {
		res.TaskResults = buildins.MakeTaskInstruction(partition, task)
	}

	res.EndTime = timestamppb.Now()
	c.ProtoBuf(http.StatusOK, res)
}
