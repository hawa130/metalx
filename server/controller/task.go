package controller

import (
	"context"

	"github.com/hawa130/metalx/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListenTask(_ *pb.Empty, stream grpc.ServerStreamingServer[pb.Task]) error {
	return status.Errorf(codes.Unimplemented, "method ListenTask not implemented")
}

func (s *Server) NotifyTaskResult(_ context.Context, result *pb.TaskResult) (*pb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyTaskResult not implemented")
}
