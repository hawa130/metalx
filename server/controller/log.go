package controller

import (
	"fmt"
	"io"

	"github.com/bwmarrin/snowflake"
	"github.com/hawa130/metalx/pb"
	"google.golang.org/grpc"
)

func (s *Server) StreamTaskLog(stream grpc.ClientStreamingServer[pb.Log, pb.Empty]) error {
	for {
		data, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Empty{})
		}
		if err != nil {
			return err
		}

		fmt.Printf("[%s] %s",
			snowflake.ParseInt64(data.TaskId).Base32(),
			data.Message,
		)
	}
}
