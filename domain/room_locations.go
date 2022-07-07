package domain

import (
	"context"
)

type RoomLocations struct {
	Model
	RoomLocationName string `json:"room_location_name"`
	RoomLocationDesc string `json:"room_location_desc"`
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
