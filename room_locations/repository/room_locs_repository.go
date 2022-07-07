package repository

import (
	"context"

	"github.com/Alterra-DataOn-Kelompok-5/room-service/domain"
	"gorm.io/gorm"
)

type mysqlRoomLocationsRepository struct {
	Db *gorm.DB
}

func NewMysqlRoomLocationsRepository(conn *gorm.DB) domain.RoomLocationsRepository {
	return &mysqlRoomLocationsRepository{conn}
}

func (m *mysqlRoomLocationsRepository) FetchAll(ctx context.Context) (res []domain.RoomLocations, err error) {
	var roomLocations []domain.RoomLocations
	err = m.Db.WithContext(ctx).Model(&domain.RoomLocations{}).Find(&roomLocations).Error

	return roomLocations, err
}

func (m *mysqlRoomLocationsRepository) FetchByID(ctx context.Context, id int) (res domain.RoomLocations, err error) {
	var roomLocation domain.RoomLocations
	err = m.Db.WithContext(ctx).Model(&domain.RoomLocations{}).Where("id = ?", id).First(&roomLocation).Error
	return roomLocation, err
}

func (m *mysqlRoomLocationsRepository) Store(ctx context.Context, rl *domain.RoomLocations) error {
	err := m.Db.WithContext(ctx).Model(&domain.RoomLocations{}).Create(&rl).Error

	return err
}

func (m *mysqlRoomLocationsRepository) Update(ctx context.Context, rl *domain.RoomLocations, id int) error {
	err := m.Db.WithContext(ctx).Model(&domain.RoomLocations{}).Where("id = ?", id).Updates(&rl).Error

	return err
}

func (m *mysqlRoomLocationsRepository) Delete(ctx context.Context, id int) error {
	err := m.Db.WithContext(ctx).Where("id = ?", id).Delete(&domain.RoomLocations{}).Error
	return err
}
