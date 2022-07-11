package repository

import (
	"context"
	"strings"

	"github.com/Alterra-DataOn-Kelompok-5/room-service/internal/dto"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/internal/model"
	pkgdto "github.com/Alterra-DataOn-Kelompok-5/room-service/pkg/dto"
	"gorm.io/gorm"
)

type RoomLocations interface {
	FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, pagination *pkgdto.Pagination) ([]model.RoomLocations, *pkgdto.PaginationInfo, error)
	FindByID(ctx context.Context, id uint) (model.RoomLocations, error)
	Save(ctx context.Context, roomLocation *dto.CreateRoomLocationsRequestBody) (model.RoomLocations, error)
	Edit(ctx context.Context, oldRoomLocation *model.RoomLocations, updateData *dto.UpdateRoomLocationsRequestBody) (*model.RoomLocations, error)
	Destroy(ctx context.Context, roomLocation *model.RoomLocations) (*model.RoomLocations, error)
	ExistByName(ctx context.Context, name string) (bool, error)
	ExistByID(ctx context.Context, id uint) (bool, error)
}

type roomLocation struct {
	Db *gorm.DB
}

func NewRoomLocationsRepository(db *gorm.DB) *roomLocation {
	return &roomLocation{
		db,
	}
}

func (r *roomLocation) FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, pagination *pkgdto.Pagination) ([]model.RoomLocations, *pkgdto.PaginationInfo, error) {
	var roomLocations []model.RoomLocations
	var count int64

	query := r.Db.WithContext(ctx).Model(&model.RoomLocations{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(room_location_name) LIKE ?", search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := pkgdto.GetLimitOffset(pagination)

	err := query.Limit(limit).Offset(offset).Find(&roomLocations).Error

	return roomLocations, pkgdto.CheckInfoPagination(pagination, count), err
}

func (r *roomLocation) FindByID(ctx context.Context, id uint) (model.RoomLocations, error) {
	var roomLocation model.RoomLocations
	if err := r.Db.WithContext(ctx).Model(&model.RoomLocations{}).Where("id = ?", id).First(&roomLocation).Error; err != nil {
		return roomLocation, err
	}
	return roomLocation, nil
}

func (r *roomLocation) Save(ctx context.Context, roomLocation *dto.CreateRoomLocationsRequestBody) (model.RoomLocations, error) {
	newRoomLocation := model.RoomLocations{
		RoomLocationName: *roomLocation.RoomLocationName,
		RoomLocationDesc: *roomLocation.RoomLocationDesc,
	}
	if err := r.Db.WithContext(ctx).Save(&newRoomLocation).Error; err != nil {
		return newRoomLocation, err
	}
	return newRoomLocation, nil
}

func (r *roomLocation) Edit(ctx context.Context, oldRoomLocation *model.RoomLocations, updateData *dto.UpdateRoomLocationsRequestBody) (*model.RoomLocations, error) {
	if updateData.RoomLocationName != nil {
		oldRoomLocation.RoomLocationName = *updateData.RoomLocationName
	}
	
	if updateData.RoomLocationDesc != nil {
		oldRoomLocation.RoomLocationDesc = *updateData.RoomLocationDesc
	}

	if err := r.Db.WithContext(ctx).Save(oldRoomLocation).Find(oldRoomLocation).Error; err != nil {
		return nil, err
	}

	return oldRoomLocation, nil
}

func (r *roomLocation) Destroy(ctx context.Context, roomLocation *model.RoomLocations) (*model.RoomLocations, error) {
	if err := r.Db.WithContext(ctx).Delete(roomLocation).Error; err != nil {
		return nil, err
	}
	return roomLocation, nil
}

func (r *roomLocation) ExistByName(ctx context.Context, name string) (bool, error) {
	var (
		count   int64
		isExist bool
	)
	if err := r.Db.WithContext(ctx).Model(&model.RoomLocations{}).Where("room_location_name = ?", name).Count(&count).Error; err != nil {
		return isExist, err
	}
	if count > 0 {
		isExist = true
	}
	return isExist, nil
}

func (r *roomLocation) ExistByID(ctx context.Context, id uint) (bool, error) {
	var (
		count   int64
		isExist bool
	)
	if err := r.Db.WithContext(ctx).Model(&model.RoomLocations{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return isExist, err
	}
	if count > 0 {
		isExist = true
	}
	return isExist, nil
}
