package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ride4Low/contracts/env"
	pb "github.com/ride4Low/contracts/proto/driver"
	"github.com/ride4Low/driver-service/internal/config"
	grpcServer "github.com/ride4Low/driver-service/internal/interface/grpc"

	"google.golang.org/grpc"
)

var (
	grpcAddr = env.GetInt("GRPC_ADDR", 9092)
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize dependency container
	container, err := config.NewContainer(ctx)
	if err != nil {
		log.Fatalf("Failed to create container: %v", err)
	}
	// defer container.Close(ctx)

	// Create and start gRPC server
	server := grpcServer.NewServer(grpcAddr)

	// Register driver service
	server.RegisterService(func(s *grpc.Server) {
		pb.RegisterDriverServiceServer(s, container.DriverHandler)
	})

	// Handle graceful shutdown
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		<-sigChan

		log.Println("Shutting down gracefully...")
		server.Stop()
		cancel()
	}()

	// Start server
	log.Println("Starting driver service...")
	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
