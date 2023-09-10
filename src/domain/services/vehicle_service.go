package services

import (
	"context"
	"time"

	"github.com/BaxtiyorBohodirov/vehicle-domain/src/domain/models"
	"github.com/BaxtiyorBohodirov/vehicle-domain/src/domain/repositories"
)

type VehicleService interface {
	CreateVehicle(ctx context.Context, driverID, model, make string) (*models.Vehicle, error)
	UpdateVehicle(ctx context.Context, vehicle *models.Vehicle) error
	DeleteVehicle(ctx context.Context, vehicleID string) error
	GetVehicle(ctx context.Context, vehicleID string) (*models.Vehicle, error)
	ListVehicleByDriver(ctx context.Context, driverID string) ([]*models.Vehicle, error)
}

type vehicleServiceImpl struct {
	vehicleRepo repositories.VehicleRepository
}

func NewVehicleService(repo *repositories.VehicleRepository) VehicleService {
	return &vehicleServiceImpl{
		vehicleRepo: *repo,
	}
}

func (v *vehicleServiceImpl) CreateVehicle(ctx context.Context, driverID, model, make string) (*models.Vehicle, error) {

	var vehicle = &models.Vehicle{
		DriverID:  driverID,
		Model:     model,
		Make:      make,
		CreatedAt: time.Now().UTC(),
	}

	err := v.vehicleRepo.SaveVehicle(ctx, vehicle)

	if err != nil {
		return nil, err
	}

	return vehicle, nil

}
func (v *vehicleServiceImpl) UpdateVehicle(ctx context.Context, vehicle *models.Vehicle) error {

	err := v.vehicleRepo.UpdateVehicle(ctx, vehicle)

	if err != nil {
		return err
	}

	return nil

}
func (v *vehicleServiceImpl) DeleteVehicle(ctx context.Context, vehicleID string) error {

	err := v.vehicleRepo.DeleteVehicle(ctx, vehicleID)

	if err != nil {
		return err
	}

	return nil
}
func (v *vehicleServiceImpl) GetVehicle(ctx context.Context, vehicleID string) (*models.Vehicle, error) {

	vehicle, err := v.vehicleRepo.GetVehicle(ctx, vehicleID)

	if err != nil {
		return nil, err
	}

	return vehicle, nil

}
func (v *vehicleServiceImpl) ListVehicleByDriver(ctx context.Context, driverID string) ([]*models.Vehicle, error) {

	vehicles, err := v.vehicleRepo.ListVehicleByDriver(ctx, driverID)

	if err != nil {
		return nil, err
	}

	return vehicles, nil
}
