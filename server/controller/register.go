package controller

import (
	"context"

	"github.com/google/uuid"
	"github.com/hawa130/metalx/pb"
)

// RegisterAgent called when edge agent becomes online
func (s *Server) RegisterAgent(ctx context.Context, in *pb.AgentId) (*pb.Empty, error) {
	id, err := uuid.FromBytes(in.Id)
	if err != nil {
		return nil, ErrInvalidAgentID
	}
	err = s.createSession(ctx, &SessionData{id: id})
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

// UnregisterAgent called when edge agent becomes offline
func (s *Server) UnregisterAgent(ctx context.Context, _ *pb.Empty) (*pb.Empty, error) {
	err := s.deleteSession(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
