package dto

import (
	"time"

	"gorm.io/gorm"
)

type (
	CreateRoomLocationsRequestBody struct {
		RoomLocationName *string `json:"room_location_name" validate:"required"`
		RoomLocationDesc *string `json:"room_location_desc" validate:"required"`
	}
	UpdateRoomLocationsRequestBody struct {
		ID               *uint   `param:"id" validate:"required"`
		RoomLocationName *string `json:"room_location_name" validate:"required"`
		RoomLocationDesc *string `json:"room_location_desc" validate:"required"`
	}
	RoomLocationsResponse struct {
		ID               uint   `json:"id"`
		RoomLocationName string `json:"room_location_name"`
		RoomLocationDesc string `json:"room_location_desc"`
	}
	RoomLocationsWithCUDResponse struct {
		RoomLocationsResponse
		CreatedAt time.Time       `json:"created_at"`
		UpdatedAt time.Time       `json:"updated_at"`
		DeletedAt *gorm.DeletedAt `json:"deleted_at"`
	}
)
