package domain

import (
	"context"
)

type Rooms struct {
	Model
	RoomName       string        `json:"room_name"`
	RoomDesc       string        `json:"room_desc"`
	RoomTypeID     int           `json:"room_type_id"`
	RoomTypes      RoomTypes     `json:"room_types" gorm:"foreignKey:RoomTypeID;references:ID"`
	RoomLocationID int           `json:"room_location_id"`
	RoomLocation   RoomLocations `json:"room_locations" gorm:"foreignKey:RoomLocationID;references:ID"`
}

type RoomsUsecase interface {
	FetchAll(ctx context.Context) ([]Rooms, error)
	FetchByID(ctx context.Context, id int) (Rooms, error)
	Store(ctx context.Context, r *Rooms) error
	Update(ctx context.Context, r *Rooms, id int) error
	Delete(ctx context.Context, id int) error
}

type RoomsRepository interface {
	FetchAll(ctx context.Context) (res []Rooms, err error)
	FetchByID(ctx context.Context, id int) (Rooms, error)
	Store(ctx context.Context, r *Rooms) error
	Update(ctx context.Context, r *Rooms, id int) error
	Delete(ctx context.Context, id int) error
}
