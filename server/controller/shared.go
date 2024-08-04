package controller

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidAgentID = status.Errorf(codes.InvalidArgument, "invalid agent id")
)
