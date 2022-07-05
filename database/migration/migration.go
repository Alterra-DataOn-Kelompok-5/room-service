package migration

import (
	"github.com/Alterra-DataOn-Kelompok-5/room-service/database"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/domain"
)

var tables = []interface{}{
	&domain.RoomTypes{},
	&domain.Rooms{},
	&domain.RoomLocations{},
	&domain.RoomDescriptions{},
}

func Migrate() {
	conn := database.GetConnection()
	conn.AutoMigrate(tables...)
}
