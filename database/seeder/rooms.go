package seeder

import (
	"log"
	"time"

	"github.com/Alterra-DataOn-Kelompok-5/room-service/internal/model"
	"gorm.io/gorm"
)

func roomsSeeder(db *gorm.DB) {
	now := time.Now()
	var rooms = []model.Rooms{
		{
			Common:         model.Common{ID: 1, CreatedAt: now, UpdatedAt: now},
			RoomName:       "Melati",
			RoomDesc:       "Ruang Meeting Melati",
			RoomTypeID:     1,
			RoomLocationID: 1,
		},
		{
			Common:         model.Common{ID: 2, CreatedAt: now, UpdatedAt: now},
			RoomName:       "Anggrek",
			RoomDesc:       "Ruang Meeting Anggrek",
			RoomTypeID:     2,
			RoomLocationID: 1,
		},
		{
			Common:         model.Common{ID: 3, CreatedAt: now, UpdatedAt: now},
			RoomName:       "Amarilis",
			RoomDesc:       "Ruang Meeting Amarilis",
			RoomTypeID:     2,
			RoomLocationID: 2,
		},
	}
	if err := db.Create(&rooms).Error; err != nil {
		log.Printf("cannot seed data rooms, with error %v\n", err)
	}
	log.Println("success seed data rooms")
}
