package repository

import (
	"context"
	"strings"

	"github.com/Alterra-DataOn-Kelompok-5/room-service/internal/dto"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/internal/model"
	pkgdto "github.com/Alterra-DataOn-Kelompok-5/room-service/pkg/dto"
	"gorm.io/gorm"
)

type Rooms interface {
	FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, p *pkgdto.Pagination) ([]model.Rooms, *pkgdto.PaginationInfo, error)
	FindByID(ctx context.Context, id uint, usePreload bool) (model.Rooms, error)
	ExistByName(ctx context.Context, name string) (bool, error)
	Save(ctx context.Context, rooms *dto.CreateRoomsRequestBody) (model.Rooms, error)
	Edit(ctx context.Context, oldRooms *model.Rooms, updateData *dto.UpdateRoomsRequestBody) (*model.Rooms, error)
	Destroy(ctx context.Context, rooms *model.Rooms) (*model.Rooms, error)
}

type rooms struct {
	Db *gorm.DB
}

func NewRoomsRepository(db *gorm.DB) *rooms {
	return &rooms{
		db,
	}
}

func (r *rooms) FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, pagination *pkgdto.Pagination) ([]model.Rooms, *pkgdto.PaginationInfo, error) {
	var rooms []model.Rooms
	var count int64

	query := r.Db.WithContext(ctx).Model(&model.Rooms{}).Preload("RoomTypes").Preload("RoomLocations")

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(room_name) LIKE ?", search, search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := pkgdto.GetLimitOffset(pagination)

	err := query.Limit(limit).Offset(offset).Find(&rooms).Error

	return rooms, pkgdto.CheckInfoPagination(pagination, count), err
}

func (r *rooms) FindByID(ctx context.Context, id uint, usePreload bool) (model.Rooms, error) {
	var room model.Rooms
	q := r.Db.WithContext(ctx).Model(&model.Rooms{}).Where("id = ?", id)
	if usePreload {
		q = q.Preload("RoomTypes").Preload("RoomLocations")
	}
	err := q.First(&room).Error
	return room, err
}

func (r *rooms) ExistByName(ctx context.Context, name string) (bool, error) {
	var (
		count   int64
		isExist bool
	)
	if err := r.Db.WithContext(ctx).Model(&model.Rooms{}).Where("room_name = ?", name).Count(&count).Error; err != nil {
		return isExist, err
	}
	if count > 0 {
		isExist = true
	}
	return isExist, nil
}

func (r *rooms) Save(ctx context.Context, rooms *dto.CreateRoomsRequestBody) (model.Rooms, error) {
	newRoom := model.Rooms{
		RoomName:       *rooms.RoomName,
		RoomDesc:       *rooms.RoomDesc,
		RoomTypeID:     *rooms.RoomTypeID,
		RoomLocationID: *rooms.RoomLocationID,
	}
	if err := r.Db.WithContext(ctx).Save(&newRoom).Error; err != nil {
		return newRoom, err
	}
	return newRoom, nil
}

func (r *rooms) Edit(ctx context.Context, oldRooms *model.Rooms, updateData *dto.UpdateRoomsRequestBody) (*model.Rooms, error) {
	if updateData.RoomName != nil {
		oldRooms.RoomName = *updateData.RoomName
	}
	if updateData.RoomDesc != nil {
		oldRooms.RoomDesc = *updateData.RoomDesc
	}
	if updateData.RoomTypeID != nil {
		oldRooms.RoomTypeID = *updateData.RoomTypeID
	}
	if updateData.RoomLocationID != nil {
		oldRooms.RoomLocationID = *updateData.RoomLocationID
	}

	if err := r.Db.
		WithContext(ctx).
		Save(oldRooms).
		Preload("RoomTypes").
		Preload("RoomLocations").
		Find(oldRooms).
		Error; err != nil {
		return nil, err
	}

	return oldRooms, nil
}

func (r *rooms) Destroy(ctx context.Context, room *model.Rooms) (*model.Rooms, error) {
	if err := r.Db.WithContext(ctx).Delete(room).Error; err != nil {
		return nil, err
	}
	return room, nil
}
