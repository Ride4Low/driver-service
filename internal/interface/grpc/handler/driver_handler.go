package handler

import (
	"context"

	pb "github.com/ride4Low/contracts/proto/driver"
)

// DriverHandler implements the gRPC driver service
type DriverHandler struct {
	pb.UnimplementedDriverServiceServer
}

func NewDriverHandler() *DriverHandler {
	return &DriverHandler{}
}

func (h *DriverHandler) RegisterDriver(ctx context.Context, req *pb.RegisterDriverRequest) (*pb.RegisterDriverResponse, error) {
	return &pb.RegisterDriverResponse{}, nil
}

func (h *DriverHandler) UnregisterDriver(ctx context.Context, req *pb.RegisterDriverRequest) (*pb.RegisterDriverResponse, error) {
	return &pb.RegisterDriverResponse{}, nil
}
