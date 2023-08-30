package worker

import (
	"context"

	"github.com/benmizrahi/gobig/internal/domains"
	"github.com/benmizrahi/gobig/internal/worker/buildins"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RouteWorkerAPIServer struct{}

// Health implements WorkerAPIServer.
func (*RouteWorkerAPIServer) Health(context.Context, *Empty) (*domains.HCRes, error) {
	res := &domains.HCRes{
		Uuid: uuid.New().String(),
		Time: timestamppb.Now(),
	}
	return res, nil
}

// TasksHandler implements WorkerAPIServer.
func (*RouteWorkerAPIServer) TasksHandler(partition *domains.IPartition, s WorkerAPI_TasksHandlerServer) error {
	res := &domains.IPartitionResult{
		TaskResults: []*domains.TaskResult{},
	}
	for _, task := range partition.Tasks {
		res.TaskResults = buildins.MakeTaskInstruction(partition, task)
	}

	res.EndTime = timestamppb.Now()
	s.Send(res)

	return nil
}

// mustEmbedUnimplementedWorkerAPIServer implements WorkerAPIServer.
func (*RouteWorkerAPIServer) mustEmbedUnimplementedWorkerAPIServer() {
	panic("unimplemented")
}
