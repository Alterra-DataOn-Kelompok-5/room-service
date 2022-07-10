package factory

import (
	"github.com/Alterra-DataOn-Kelompok-5/room-service/database"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/internal/repository"
)

type Factory struct {
	RoomsRepository         repository.Rooms
	RoomLocationsRepository repository.RoomLocations
	RoomTypesRepository     repository.RoomTypes
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		repository.NewRoomsRepository(db),
		repository.NewRoomLocationsRepository(db),
		repository.NewRoomTypesRepository(db),
	}
}
