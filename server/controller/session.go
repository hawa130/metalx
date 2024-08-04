package controller

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func getSessionId(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.InvalidArgument, "Missing metadata")
	}
	sessionIds := md["session-id"]
	if len(sessionIds) == 0 {
		return "", nil
	}
	sessionId := sessionIds[0]
	return sessionId, nil
}

func (s *Server) getSession(ctx context.Context) (*SessionData, error) {
	sessionId, err := getSessionId(ctx)
	if err != nil {
		return nil, err
	}
	if sessionId == "" {
		return nil, nil
	}
	session, exists := s.sessions[sessionId]
	if !exists {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid Session")
	}
	return session, nil
}

func (s *Server) createSession(ctx context.Context, session *SessionData) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	strId := id.String()
	s.lock.Lock()
	defer s.lock.Unlock()
	if session == nil {
		return status.Errorf(codes.InvalidArgument, "Invalid Session Initialization")
	}
	s.sessions[strId] = session
	header := metadata.Pairs("set-session-id", strId)
	err = grpc.SendHeader(ctx, header)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) deleteSession(ctx context.Context) error {
	sessionId, err := getSessionId(ctx)
	if err != nil {
		return err
	}
	if sessionId == "" {
		return status.Errorf(codes.InvalidArgument, "Invalid Session")
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.sessions, sessionId)
	header := metadata.Pairs("set-session-id", "")
	err = grpc.SendHeader(ctx, header)
	if err != nil {
		return err
	}
	return nil
}
