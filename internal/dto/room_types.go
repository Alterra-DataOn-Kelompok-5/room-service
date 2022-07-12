package dto

import (
	"time"

	"gorm.io/gorm"
)

type (
	CreateRoomTypesRequestBody struct {
		RoomTypeName        *string `json:"room_type_name" validate:"required"`
		RoomTypeMaxCapacity *int    `json:"room_type_max_capacity" validate:"required"`
		RoomTypeDesc        *string `json:"room_type_desc" validate:"required"`
	}
	UpdateRoomTypesRequestBody struct {
		ID                  *uint   `param:"id" validate:"required"`
		RoomTypeName        *string `json:"room_type_name" validate:"omitempty"`
		RoomTypeMaxCapacity *int    `json:"room_type_max_capacity" validate:"omitempty"`
		RoomTypeDesc        *string `json:"room_type_desc" validate:"omitempty"`
	}
	RoomTypesResponse struct {
		ID                  uint   `json:"id"`
		RoomTypeName        string `json:"room_type_name"`
		RoomTypeMaxCapacity int    `json:"room_type_max_capacity"`
		RoomTypeDesc        string `json:"room_type_desc"`
	}
	RoomTypesWithCUDResponse struct {
		RoomTypesResponse
		CreatedAt time.Time       `json:"created_at"`
		UpdatedAt time.Time       `json:"updated_at"`
		DeletedAt *gorm.DeletedAt `json:"deleted_at"`
	}
)
