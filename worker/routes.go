package worker

import (
	"io"
	"log"
	"net/http"

	"github.com/benmizrahi/godist/protos"
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

func (w *Worker) tasksHandler(c *gin.Context) {
	buf, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalln("Failed to parse register request:", err)
	}

	req := &protos.IPartition{}
	if err := proto.Unmarshal(buf, req); err != nil {
		log.Fatalln("Failed to parse register request:", err)
	}

	//TODO: work
	//

	res := protos.IPartitionResult{
		TaskResults: []*protos.TaskResult{},
	}

	for _, task := range req.Tasks {
		res.TaskResults = append(res.TaskResults, &protos.TaskResult{
			Uuid:    task.Uuid,
			Status:  true,
			EndTime: timestamppb.Now(),
		})
	}

	res.EndTime = timestamppb.Now()
	c.ProtoBuf(http.StatusOK, res)
}
