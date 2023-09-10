package vehicle

import (
	"context"

	"github.com/BaxtiyorBohodirov/vehicle-domain/src/domain/models"
	"gorm.io/gorm"
)

const (
	vehicleTable = "vehicle.vehicles"
)

type vehicleRepositoryImpl struct {
	db gorm.DB
}

func NewVehicleRepository(db gorm.DB) *vehicleRepositoryImpl {
	return &vehicleRepositoryImpl{
		db: db,
	}
}

func (v *vehicleRepositoryImpl) SaveVehicle(ctx context.Context, vehicle *models.Vehicle) error {

	result := v.db.WithContext(ctx).Table(vehicleTable).Create(&vehicle)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
func (v *vehicleRepositoryImpl) UpdateVehicle(ctx context.Context, vehicle *models.Vehicle) error {

	result := v.db.WithContext(ctx).Table(vehicleTable).Save(&vehicle)

	if result.Error != nil {
		return result.Error
	}

	return nil

}
func (v *vehicleRepositoryImpl) DeleteVehicle(ctx context.Context, vehicleID string) error {
	result := v.db.WithContext(ctx).Delete(&models.Vehicle{}, vehicleID)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (v *vehicleRepositoryImpl) GetVehicle(ctx context.Context, vehicleID string) (*models.Vehicle, error) {
	var vehicle *models.Vehicle

	result := v.db.WithContext(ctx).Where("id = ?", vehicleID).First(&vehicle)

	if result.Error != nil {
		return nil, result.Error
	}

	return vehicle, nil
}

func (v *vehicleRepositoryImpl) ListVehicleByDriver(ctx context.Context, driverID string) ([]*models.Vehicle, error) {
	var vehicles []*models.Vehicle

	result := v.db.WithContext(ctx).Table(vehicleTable).Where("driver_id = ?", driverID).Find(&vehicles)

	if result.Error != nil {
		return nil, result.Error
	}

	return vehicles, nil
}
