package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/ride4Low/contracts/env"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server represents the gRPC server
type Server struct {
	server *grpc.Server
	port   int
}

// NewServer creates a new gRPC server
func NewServer(port int) *Server {
	return &Server{
		server: grpc.NewServer(),
		port:   port,
	}
}

// RegisterService registers a gRPC service
func (s *Server) RegisterService(registerFunc func(*grpc.Server)) {
	registerFunc(s.server)
}

// Start starts the gRPC server
func (s *Server) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	if env.GetString("APP_ENV", "development") == "development" {
		// Register reflection service for development
		reflection.Register(s.server)
	}

	log.Printf("gRPC server listening on port %d", s.port)

	if err := s.server.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}

// Stop stops the gRPC server gracefully
func (s *Server) Stop() {
	log.Println("Stopping gRPC server...")
	s.server.GracefulStop()
}
