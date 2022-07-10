package seeder

import (
	"log"
	"time"

	"github.com/Alterra-DataOn-Kelompok-5/room-service/internal/model"
	"gorm.io/gorm"
)

func roomTypesSeeder(db *gorm.DB) {
	now := time.Now()
	var roomTypes = []model.RoomTypes{
		{
			Common: model.Common{
				ID:        1,
				CreatedAt: now,
				UpdatedAt: now,
			},
			RoomTypeName:        "Small Meeting Room",
			RoomTypeMaxCapacity: 5,
			RoomTypeDesc:        "Ruang meeting kecil dengan kapasitas 1-5 orang",
		},
		{
			Common: model.Common{
				ID:        2,
				CreatedAt: now,
				UpdatedAt: now,
			},
			RoomTypeName:        "Medium Meeting Room",
			RoomTypeMaxCapacity: 10,
			RoomTypeDesc:        "Ruang meeting sedang dengan kapasitas 6-10 orang",
		},
		{
			Common: model.Common{
				ID:        3,
				CreatedAt: now,
				UpdatedAt: now,
			},
			RoomTypeName:        "Large Meeting Room",
			RoomTypeMaxCapacity: 20,
			RoomTypeDesc:        "Ruang meeting kecil dengan kapasitas 11-20 orang",
		},
	}
	if err := db.Create(&roomTypes).Error; err != nil {
		log.Printf("cannot seed data room types, with error %v\n", err)
	}
	log.Println("success seed data room types")
}
