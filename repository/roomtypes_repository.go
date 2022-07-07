package repository

import (
	"context"

	"github.com/Alterra-DataOn-Kelompok-5/room-service/domain"
	"gorm.io/gorm"
)

type mysqlRoomTypesRepository struct {
	Db *gorm.DB
}

func NewMysqlRoomTypesRepository(conn *gorm.DB) domain.RoomTypesRepository {
	return &mysqlRoomTypesRepository{conn}
}

func (m *mysqlRoomTypesRepository) FetchAll(ctx context.Context) (res []domain.RoomTypes, err error) {
	var roomTypes []domain.RoomTypes
	err = m.Db.WithContext(ctx).Model(&domain.RoomTypes{}).Find(&roomTypes).Error

	return roomTypes, err
}

func (m *mysqlRoomTypesRepository) FetchByID(ctx context.Context, id int) (res domain.RoomTypes, err error) {
	var roomType domain.RoomTypes
	err = m.Db.WithContext(ctx).Model(&domain.RoomTypes{}).Where("id = ?", id).First(&roomType).Error
	return roomType, err
}

func (m *mysqlRoomTypesRepository) Store(ctx context.Context, rt *domain.RoomTypes) error {
	err := m.Db.WithContext(ctx).Model(&domain.RoomTypes{}).Create(&rt).Error

	return err
}

func (m *mysqlRoomTypesRepository) Update(ctx context.Context, rt *domain.RoomTypes, id int) error {
	err := m.Db.WithContext(ctx).Model(&domain.RoomTypes{}).Where("id = ?", id).Updates(&rt).Error

	return err
}

func (m *mysqlRoomTypesRepository) Delete(ctx context.Context, id int) error {
	err := m.Db.WithContext(ctx).Where("id = ?", id).Delete(&domain.RoomTypes{}).Error
	return err
}
