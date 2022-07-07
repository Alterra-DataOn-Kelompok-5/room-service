package repository

import (
	"context"

	"github.com/Alterra-DataOn-Kelompok-5/room-service/domain"
	"gorm.io/gorm"
)

type mysqlRoomsRepository struct {
	Db *gorm.DB
}

func NewMysqlRoomsRepository(conn *gorm.DB) domain.RoomsRepository {
	return &mysqlRoomsRepository{conn}
}

func (m *mysqlRoomsRepository) FetchAll(ctx context.Context) (res []domain.Rooms, err error) {
	var rooms []domain.Rooms
	err = m.Db.WithContext(ctx).Model(&domain.Rooms{}).Find(&rooms).Error

	return rooms, err
}

func (m *mysqlRoomsRepository) FetchByID(ctx context.Context, id int) (res domain.Rooms, err error) {
	var room domain.Rooms
	err = m.Db.WithContext(ctx).Model(&domain.Rooms{}).Where("id = ?", id).First(&room).Error
	return room, err
}

func (m *mysqlRoomsRepository) Store(ctx context.Context, r *domain.Rooms) error {
	err := m.Db.WithContext(ctx).Model(&domain.Rooms{}).Create(&r).Error

	return err
}

func (m *mysqlRoomsRepository) Update(ctx context.Context, r *domain.Rooms, id int) error {
	err := m.Db.WithContext(ctx).Model(&domain.Rooms{}).Where("id = ?", id).Updates(&r).Error

	return err
}

func (m *mysqlRoomsRepository) Delete(ctx context.Context, id int) error {
	err := m.Db.WithContext(ctx).Where("id = ?", id).Delete(&domain.Rooms{}).Error
	return err
}