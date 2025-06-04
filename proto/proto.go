package proto

import (
	"context"

	"github.com/nktauserum/crawler-service/pkg/db"
	"github.com/nktauserum/crawler-service/pkg/task"
	"github.com/nktauserum/crawler-service/proto/pb"
)

type Server struct {
	pb.TaskServiceServer
}

func (s *Server) GetAvailableTask(context.Context, *pb.Empty) (*pb.Task, error) {
	st := db.GetStorage()
	task, err := task.PendingTask(st)
	if err != nil {
		return nil, err
	}

	return &pb.Task{
		Uuid: task.UUID,
		Url:  task.URL,
	}, nil
}

func (s *Server) CompleteTask(ctx context.Context, taskResult *pb.TaskResult) (*pb.Empty, error) {
	st := db.GetStorage()

	err := task.CompleteTask(st, taskResult.Uuid, taskResult.Result)

	return &pb.Empty{}, err
}
