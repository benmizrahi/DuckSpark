package worker

import (
	"io"
	"log"
	"net/http"

	"github.com/benmizrahi/gobig/internal/protos"
	"github.com/benmizrahi/gobig/internal/worker/buildins"
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

	plan := &protos.TasksPlan{}
	if err := proto.Unmarshal(buf, plan); err != nil {
		log.Fatalln("Failed to parse register request:", err)
	}

	res := []*protos.TaskResult{}
	for _, task := range plan.Tasks {
		if task.Plugin != "" {
			res = append(res, w.Plugins[task.Plugin].Execute(task))
		}
		res = append(res, buildins.MakeInstactions(task))
	}
	c.ProtoBuf(http.StatusOK, res)
}
