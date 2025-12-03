package repository

import (
	"context"

	pb "github.com/ride4Low/contracts/proto/driver"
)

type DriverRepository interface {
	Create(ctx context.Context, driver *pb.Driver) error
	Remove(ctx context.Context, driverId string) error
	GetIDByPackageSlug(ctx context.Context, packageSlug string) ([]string, error)
}
