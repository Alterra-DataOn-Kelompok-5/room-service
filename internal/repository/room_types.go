package repository

import (
	"context"
	"strings"

	"github.com/Alterra-DataOn-Kelompok-5/room-service/internal/dto"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/internal/model"
	pkgdto "github.com/Alterra-DataOn-Kelompok-5/room-service/pkg/dto"
	"gorm.io/gorm"
)

type RoomTypes interface {
	FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, p *pkgdto.Pagination) ([]model.RoomTypes, *pkgdto.PaginationInfo, error)
	FindByID(ctx context.Context, id uint) (model.RoomTypes, error)
	Save(ctx context.Context, roomTypes *dto.CreateRoomTypesRequestBody) (model.RoomTypes, error)
	Edit(ctx context.Context, oldRoomTypes *model.RoomTypes, updateData *dto.UpdateRoomTypesRequestBody) (*model.RoomTypes, error)
	Destroy(ctx context.Context, roomTypes *model.RoomTypes) (*model.RoomTypes, error)
	ExistByName(ctx context.Context, name string) (bool, error)
}

type roomTypes struct {
	Db *gorm.DB
}

func NewRoomTypesRepository(db *gorm.DB) *roomTypes {
	return &roomTypes{
		db,
	}
}

func (r *roomTypes) FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, pagination *pkgdto.Pagination) ([]model.RoomTypes, *pkgdto.PaginationInfo, error) {
	var roomTypes []model.RoomTypes
	var count int64

	query := r.Db.WithContext(ctx).Model(&model.RoomTypes{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(room_type_name) LIKE ?", search, search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := pkgdto.GetLimitOffset(pagination)

	err := query.Limit(limit).Offset(offset).Find(&roomTypes).Error

	return roomTypes, pkgdto.CheckInfoPagination(pagination, count), err
}

func (r *roomTypes) FindByID(ctx context.Context, id uint) (model.RoomTypes, error) {
	var roomType model.RoomTypes
	if err := r.Db.WithContext(ctx).Model(&model.RoomTypes{}).Where("id = ?", id).First(&roomType).Error; err != nil {
		return roomType, err
	}
	return roomType, nil
}

func (r *roomTypes) Save(ctx context.Context, roomType *dto.CreateRoomTypesRequestBody) (model.RoomTypes, error) {
	newRoomType := model.RoomTypes{
		RoomTypeName:        *roomType.RoomTypeName,
		RoomTypeMaxCapacity: *roomType.RoomTypeMaxCapacity,
		RoomTypeDesc:        *roomType.RoomTypeDesc,
	}
	if err := r.Db.WithContext(ctx).Save(&newRoomType).Error; err != nil {
		return newRoomType, err
	}
	return newRoomType, nil
}

func (r *roomTypes) Edit(ctx context.Context, oldRoomTypes *model.RoomTypes, updateData *dto.UpdateRoomTypesRequestBody) (*model.RoomTypes, error) {
	if updateData.RoomTypeName != nil {
		oldRoomTypes.RoomTypeName = *updateData.RoomTypeName
		oldRoomTypes.RoomTypeMaxCapacity = *updateData.RoomTypeMaxCapacity
		oldRoomTypes.RoomTypeDesc = *updateData.RoomTypeDesc
	}

	if err := r.Db.WithContext(ctx).Save(oldRoomTypes).Find(oldRoomTypes).Error; err != nil {
		return nil, err
	}

	return oldRoomTypes, nil
}

func (r *roomTypes) Destroy(ctx context.Context, roomType *model.RoomTypes) (*model.RoomTypes, error) {
	if err := r.Db.WithContext(ctx).Delete(roomType).Error; err != nil {
		return nil, err
	}
	return roomType, nil
}

func (r *roomTypes) ExistByName(ctx context.Context, name string) (bool, error) {
	var (
		count   int64
		isExist bool
	)
	if err := r.Db.WithContext(ctx).Model(&model.RoomTypes{}).Where("room_type_name = ?", name).Count(&count).Error; err != nil {
		return isExist, err
	}
	if count > 0 {
		isExist = true
	}
	return isExist, nil
}
