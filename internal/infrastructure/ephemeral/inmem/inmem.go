package inmem

import (
	"context"
	"sync"

	pb "github.com/ride4Low/contracts/proto/driver"
	"github.com/ride4Low/driver-service/internal/domain/repository"
)

type driverRepository struct {
	drivers map[string]*pb.Driver

	mu sync.RWMutex
}

func NewDriverRepository() repository.DriverRepository {
	return &driverRepository{
		drivers: make(map[string]*pb.Driver),
		mu:      sync.RWMutex{},
	}
}

// RegisterDriver means driver is waiting for rider
func (r *driverRepository) Create(ctx context.Context, driver *pb.Driver) (*pb.Driver, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.drivers[driver.Id] = driver

	return driver, nil
}

// UnregisterDriver means driver is not waiting for rider
func (r *driverRepository) Remove(ctx context.Context, driverId string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.drivers, driverId)

	return nil
}

func (r *driverRepository) GetIDsByPackageSlug(ctx context.Context, packageSlug string) ([]string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var driverIds []string
	for _, driver := range r.drivers {
		if driver.PackageSlug == packageSlug {
			driverIds = append(driverIds, driver.Id)
		}
	}

	return driverIds, nil
}
