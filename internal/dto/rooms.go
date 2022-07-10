package dto

import (
	"time"

	"gorm.io/gorm"
)

type (
	CreateRoomsRequestBody struct {
		RoomName       *string `json:"room_name" validate:"required"`
		RoomDesc       *string `json:"room_desc" validate:"omitempty"`
		RoomTypeID     *int    `json:"room_type_id" validate:"omitempty"`
		RoomLocationID *int    `json:"room_location_id" validate:"omitempty"`
	}
	UpdateRoomsRequestBody struct {
		ID             *uint   `param:"id" validate:"required"`
		RoomName       *string `json:"room_name" validate:"omitempty"`
		RoomDesc       *string `json:"room_desc" validate:"omitempty"`
		RoomTypeID     *int    `json:"room_type_id" validate:"omitempty"`
		RoomLocationID *int    `json:"room_location_id" validate:"omitempty"`
	}
	RoomsResponse struct {
		ID       uint   `json:"id"`
		RoomName string `json:"room_name"`
		RoomDesc string `json:"room_desc"`
	}
	RoomsWithCUDResponse struct {
		RoomsResponse
		CreatedAt time.Time       `json:"created_at"`
		UpdatedAt time.Time       `json:"updated_at"`
		DeletedAt *gorm.DeletedAt `json:"deleted_at"`
	}
	RoomsDetailResponse struct {
		RoomsResponse
		RoomTypes     RoomTypesResponse     `json:"room_types"`
		RoomLocations RoomLocationsResponse `json:"room_locations"`
	}
)
