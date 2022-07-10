package seeder

import (
	"github.com/Alterra-DataOn-Kelompok-5/room-service/database"
	"gorm.io/gorm"
)

type seed struct {
	DB *gorm.DB
}

func NewSeeder() *seed {
	return &seed{database.GetConnection()}
}

func (s *seed) SeedAll() {
	roomTypesSeeder(s.DB)
	roomLocationsSeeder(s.DB)
	roomsSeeder(s.DB)
}

func (s *seed) DeleteAll() {
	s.DB.Exec("DELETE FROM rooms")
	s.DB.Exec("DELETE FROM room_types")
	s.DB.Exec("DELETE FROM room_locations")
}
