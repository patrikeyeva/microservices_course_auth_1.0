package server

import (
	"context"
	"log"

	"github.com/patrikeyeva/microservices_course_auth_1.0/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Server implements gRPC auth server
type Server struct {
	user_v1.UnimplementedUserV1Server
}

// Create creating new user
func (s *Server) Create(_ context.Context, req *user_v1.CreateRequest) (*user_v1.CreateResponse, error) {
	log.Printf("CREATE:\n %v", req.GetData())
	return &user_v1.CreateResponse{
		Id: 0,
	}, nil
}

// Get getting information about a user by his ID
func (s *Server) Get(_ context.Context, req *user_v1.GetRequest) (*user_v1.GetResponse, error) {
	log.Printf("GET:\n %v", req.GetId())
	return &user_v1.GetResponse{
		Data: &user_v1.UserGet{
			Id:        req.GetId(),
			Name:      "",
			Email:     "",
			Role:      0,
			CreatedAt: &timestamppb.Timestamp{},
			UpdatedAt: &timestamppb.Timestamp{},
		},
	}, nil
}

// Update updating user information by user ID
func (s *Server) Update(_ context.Context, req *user_v1.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("UPDATE:\n %v", req.GetData())
	return nil, nil
}

// Delete deleting user by ID
func (s *Server) Delete(_ context.Context, req *user_v1.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("DELETE:\n %v", req.GetId())
	return nil, nil
}
