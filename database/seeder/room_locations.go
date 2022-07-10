package seeder

import (
	"log"
	"time"

	"github.com/Alterra-DataOn-Kelompok-5/room-service/internal/model"
	"gorm.io/gorm"
)

func roomLocationsSeeder(db *gorm.DB) {
	now := time.Now()
	var roomLocation = []model.RoomLocations{
		{
			Common: model.Common{
				ID:        1,
				CreatedAt: now,
				UpdatedAt: now,
			},
			RoomLocationName: "1F",
			RoomLocationDesc: "Lantai 1",
		},
		{
			Common: model.Common{
				ID:        2,
				CreatedAt: now,
				UpdatedAt: now,
			},
			RoomLocationName: "2F Right Wing",
			RoomLocationDesc: "Lantai 2 bagian kanan gedung",
		},
		{
			Common: model.Common{
				ID:        3,
				CreatedAt: now,
				UpdatedAt: now,
			},
			RoomLocationName: "2F Left Wing",
			RoomLocationDesc: "Lantai 2 bagian kiri gedung",
		},
	}
	if err := db.Create(&roomLocation).Error; err != nil {
		log.Printf("cannot seed data room locations, with error %v\n", err)
	}
	log.Println("success seed data room locations")
}
