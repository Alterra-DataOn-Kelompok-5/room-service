package migration

import (
	"github.com/Alterra-DataOn-Kelompok-5/room-service/database"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/domain"
)

var tables = []interface{}{
	&domain.RoomTypes{},
	&domain.RoomLocations{},
	&domain.Rooms{},
}

func Migrate() {
	conn := database.GetConnection()
	conn.AutoMigrate(tables...)
}
