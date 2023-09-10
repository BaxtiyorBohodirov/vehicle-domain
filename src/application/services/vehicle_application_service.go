package services

import (
	"context"
	"errors"

	"github.com/BaxtiyorBohodirov/vehicle-domain/src/application/dtos"
	"github.com/BaxtiyorBohodirov/vehicle-domain/src/domain/models"
	vehicleServices "github.com/BaxtiyorBohodirov/vehicle-domain/src/domain/services"
)

type VehicleApplicationService interface {
	CreateVehicle(ctx context.Context, driverID, model, make string) (*dtos.GetVehicleResponse, error)
	UpdateVehicle(ctx context.Context, driverID, model, make string) (*dtos.GetVehicleResponse, error)
	DeleteVehicle(ctx context.Context, vehicleID string) error
	GetVehicle(ctx context.Context, vehicleID string) (*dtos.GetVehicleResponse, error)
	ListVehicleByDriver(ctx context.Context, driverID string) ([]*models.Vehicle, error)
}

type vehicleApplicationServiceImpl struct {
	vehicleService vehicleServices.VehicleService
}

func NewVehicleService(service *vehicleServices.VehicleService) VehicleApplicationService {
	return &vehicleApplicationServiceImpl{
		vehicleService: *service,
	}
}

func (v *vehicleApplicationServiceImpl) CreateVehicle(ctx context.Context, driverID, model, make string) (*dtos.GetVehicleResponse, error) {

	if driverID == "" {
		return nil, errors.New("Driver Id is required")
	}

	if model == "" {
		return nil, errors.New("Model is required")
	}

	if make == "" {
		return nil, errors.New("Make is required")
	}

	vehicle, err := v.vehicleService.CreateVehicle(ctx, driverID, model, make)

	if err != nil {
		return nil, err
	}

	return dtos.NewGetVehicleResponse(vehicle), nil

}
func (v *vehicleApplicationServiceImpl) UpdateVehicle(ctx context.Context, driverID, model, make string) (*dtos.GetVehicleResponse, error) {

	if driverID == "" {
		return nil, errors.New("Driver Id is required")
	}

	if model == "" {
		return nil, errors.New("Model is required")
	}

	if make == "" {
		return nil, errors.New("Make is required")
	}

	var vehicle = &models.Vehicle{
		DriverID: driverID,
		Model:    model,
		Make:     make,
	}

	err := v.vehicleService.UpdateVehicle(ctx, vehicle)

	if err != nil {
		return nil, err
	}

	return dtos.NewGetVehicleResponse(vehicle), nil

}
func (v *vehicleApplicationServiceImpl) DeleteVehicle(ctx context.Context, vehicleID string) error {

	if vehicleID == "" {
		return errors.New("Vehicle Id is required")
	}

	err := v.vehicleService.DeleteVehicle(ctx, vehicleID)

	if err != nil {
		return err
	}

	return nil
}
func (v *vehicleApplicationServiceImpl) GetVehicle(ctx context.Context, vehicleID string) (*dtos.GetVehicleResponse, error) {

	if vehicleID == "" {
		return nil, errors.New("Vehicle Id is required")
	}

	vehicle, err := v.vehicleService.GetVehicle(ctx, vehicleID)

	if err != nil {
		return nil, err
	}

	return dtos.NewGetVehicleResponse(vehicle), nil

}
func (v *vehicleApplicationServiceImpl) ListVehicleByDriver(ctx context.Context, driverID string) ([]*models.Vehicle, error) {

	if driverID == "" {
		return nil, errors.New("Driver Id is required")
	}

	vehicles, err := v.vehicleService.ListVehicleByDriver(ctx, driverID)

	if err != nil {
		return nil, err
	}

	return vehicles, nil
}
