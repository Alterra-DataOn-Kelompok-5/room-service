package domain

import (
	"context"
)

type RoomLocations struct {
	Model
	RoomLocationName string `json:"room_location"`
	RoomFloor        string `json:"room_floor"`
	RoomID           int    `json:"room_id"`
	Rooms            Rooms  `json:"rooms" gorm:"foreignKey:RoomID;references:ID"`
}

type RoomLocationsUsecase interface {
	FetchAll(ctx context.Context) ([]RoomLocations, error)
	FetchByID(ctx context.Context, id int) (RoomLocations, error)
	Store(ctx context.Context, rl *RoomLocations) error
	Update(ctx context.Context, rl *RoomLocations, id int) error
	Delete(ctx context.Context, id int) error
}

type RoomLocationsRepository interface {
	FetchAll(ctx context.Context) (res []RoomLocations, err error)
	FetchByID(ctx context.Context, id int) (RoomLocations, error)
	Store(ctx context.Context, rl *RoomLocations) error
	Update(ctx context.Context, rl *RoomLocations, id int) error
	Delete(ctx context.Context, id int) error
}
