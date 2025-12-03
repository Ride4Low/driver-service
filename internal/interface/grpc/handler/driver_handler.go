package handler

import (
	"context"

	pb "github.com/ride4Low/contracts/proto/driver"
	"github.com/ride4Low/driver-service/internal/application/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DriverHandler implements the gRPC driver service
type DriverHandler struct {
	pb.UnimplementedDriverServiceServer
	driverService service.DriverService
}

func NewDriverHandler(driverService service.DriverService) *DriverHandler {
	return &DriverHandler{
		driverService: driverService,
	}
}

func (h *DriverHandler) RegisterDriver(ctx context.Context, req *pb.RegisterDriverRequest) (*pb.RegisterDriverResponse, error) {
	driver, err := h.driverService.RegisterDriver(ctx, req.GetDriverID(), req.GetPackageSlug())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to register driver")
	}
	return &pb.RegisterDriverResponse{
		Driver: driver,
	}, nil
}

func (h *DriverHandler) UnregisterDriver(ctx context.Context, req *pb.RegisterDriverRequest) (*pb.RegisterDriverResponse, error) {
	err := h.driverService.UnregisterDriver(ctx, req.GetDriverID())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unregister driver")
	}
	return &pb.RegisterDriverResponse{
		Driver: &pb.Driver{
			Id: req.GetDriverID(),
		},
	}, nil
}
