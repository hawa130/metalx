package controller

import (
	"context"

	"github.com/hawa130/metalx/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Ping the controller, used to check if controller is online and reachable
func (s *Server) Ping(context.Context, *pb.Empty) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (s *Server) StreamEvent(stream grpc.ClientStreamingServer[pb.Event, pb.Empty]) error {
	return status.Errorf(codes.Unimplemented, "method StreamEvent not implemented")
}
