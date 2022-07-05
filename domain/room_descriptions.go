package domain

import (
	"context"
)

type RoomDescriptions struct {
	Model
	RoomDesc string `json:"room_desc"`
	RoomID   int    `json:"room_id"`
	Rooms    Rooms  `json:"rooms" gorm:"foreignKey:RoomID;references:ID"`
}

type RoomDescriptionsUsecase interface {
	FetchAll(ctx context.Context) ([]RoomDescriptions, error)
	FetchByID(ctx context.Context, id int) (RoomDescriptions, error)
	Store(ctx context.Context, rd *RoomDescriptions) error
	Update(ctx context.Context, rd *RoomDescriptions, id int) error
	Delete(ctx context.Context, id int) error
}

type RoomDescriptionsRepository interface {
	FetchAll(ctx context.Context) (res []RoomDescriptions, err error)
	FetchByID(ctx context.Context, id int) (RoomDescriptions, error)
	Store(ctx context.Context, rd *RoomDescriptions) error
	Update(ctx context.Context, rd *RoomDescriptions, id int) error
	Delete(ctx context.Context, id int) error
}
