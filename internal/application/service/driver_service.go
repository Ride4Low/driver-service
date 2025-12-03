package service

import (
	"context"
	math "math/rand/v2"

	"github.com/mmcloughlin/geohash"
	pb "github.com/ride4Low/contracts/proto/driver"
	"github.com/ride4Low/driver-service/internal/domain/repository"
)

type DriverService interface {
	RegisterDriver(ctx context.Context, driverId string, packageSlug string) error
	UnregisterDriver(ctx context.Context, driverId string) error
	FindAvailableDrivers(ctx context.Context, packageSlug string) ([]string, error)
}

type driverService struct {
	driverRepo repository.DriverRepository
}

func NewDriverService(driverRepo repository.DriverRepository) DriverService {
	return &driverService{
		driverRepo: driverRepo,
	}
}

func (s *driverService) RegisterDriver(ctx context.Context, driverId string, packageSlug string) error {
	randomIndex := math.IntN(len(PredefinedRoutes))
	randomRoute := PredefinedRoutes[randomIndex]

	randomPlate := GenerateRandomPlate()
	randomAvatar := GetRandomAvatar(randomIndex)

	// we can ignore this property for now, but it must be sent to the frontend.
	geohash := geohash.Encode(randomRoute[0][0], randomRoute[0][1])

	driver := &pb.Driver{
		Id:             driverId,
		Geohash:        geohash,
		Location:       &pb.Location{Latitude: randomRoute[0][0], Longitude: randomRoute[0][1]},
		Name:           "Vin Diesel",
		PackageSlug:    packageSlug,
		ProfilePicture: randomAvatar,
		CarPlate:       randomPlate,
	}
	return s.driverRepo.Create(ctx, driver)
}

func (s *driverService) UnregisterDriver(ctx context.Context, driverId string) error {
	return s.driverRepo.Remove(ctx, driverId)
}

func (s *driverService) FindAvailableDrivers(ctx context.Context, packageSlug string) ([]string, error) {
	return s.driverRepo.GetIDByPackageSlug(ctx, packageSlug)
}
