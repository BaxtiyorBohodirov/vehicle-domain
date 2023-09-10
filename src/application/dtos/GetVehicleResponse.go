package dtos

import "github.com/BaxtiyorBohodirov/vehicle-domain/src/domain/models"

type GetVehicleResponse struct {
	Vehicle *models.Vehicle `json:"rating"`
}

func NewGetVehicleResponse(vehicle *models.Vehicle) *GetVehicleResponse {
	return &GetVehicleResponse{
		Vehicle: vehicle,
	}
}
