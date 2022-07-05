package domain

import (
	"context"
)

type RoomTypes struct {
	Model
	RoomTypeName     string `json:"room_name"`
	RoomTypeCapacity int    `json:"room_type_capacity"`
	RoomTypeDesc     string `json:"room_type_desc" gorm:"type:text"`
}

type RoomTypesUsecase interface {
	FetchAll(ctx context.Context) ([]RoomTypes, error)
	FetchByID(ctx context.Context, id int) (RoomTypes, error)
	Store(ctx context.Context, rt *RoomTypes) error
	Update(ctx context.Context, rt *RoomTypes, id int) error
	Delete(ctx context.Context, id int) error
}

type RoomTypesRepository interface {
	FetchAll(ctx context.Context) (res []RoomTypes, err error)
	FetchByID(ctx context.Context, id int) (RoomTypes, error)
	Store(ctx context.Context, rt *RoomTypes) error
	Update(ctx context.Context, rt *RoomTypes, id int) error
	Delete(ctx context.Context, id int) error
}
