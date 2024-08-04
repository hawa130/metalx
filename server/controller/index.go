package controller

import (
	"sync"

	"github.com/google/uuid"
	"github.com/hawa130/metalx/pb"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedControllerServer
	sessions map[string]*SessionData
	lock     sync.Mutex
}

type SessionData struct {
	id uuid.UUID
}

func Register(srv *grpc.Server) {
	s := &Server{
		sessions: make(map[string]*SessionData),
	}
	pb.RegisterControllerServer(srv, s)
}
