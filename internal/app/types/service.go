package types

import (
	"context"
	"errors"

	"github.com/Alterra-DataOn-Kelompok-5/room-service/internal/dto"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/internal/factory"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/internal/repository"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/pkg/constant"
	pkgdto "github.com/Alterra-DataOn-Kelompok-5/room-service/pkg/dto"
	res "github.com/Alterra-DataOn-Kelompok-5/room-service/pkg/util/response"
)

type service struct {
	RoomTypesRepository repository.RoomTypes
}

type Service interface {
	Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto.RoomTypesResponse], error)
	FindByID(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.RoomTypesResponse, error)
	Store(ctx context.Context, payload *dto.CreateRoomTypesRequestBody) (*dto.RoomTypesResponse, error)
	UpdateById(ctx context.Context, payload *dto.UpdateRoomTypesRequestBody) (*dto.RoomTypesResponse, error)
	DeleteById(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.RoomTypesWithCUDResponse, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		RoomTypesRepository: f.RoomTypesRepository,
	}
}

func (s *service) Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto.RoomTypesResponse], error) {
	roomTypes, info, err := s.RoomTypesRepository.FindAll(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	var data []dto.RoomTypesResponse

	for _, roomType := range roomTypes {
		data = append(data, dto.RoomTypesResponse{
			ID:                  roomType.ID,
			RoomTypeName:        roomType.RoomTypeName,
			RoomTypeMaxCapacity: roomType.RoomTypeMaxCapacity,
			RoomTypeDesc:        roomType.RoomTypeDesc,
		})

	}

	result := new(pkgdto.SearchGetResponse[dto.RoomTypesResponse])
	result.Data = data
	result.PaginationInfo = *info

	return result, nil
}
func (s *service) FindByID(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.RoomTypesResponse, error) {
	var result dto.RoomTypesResponse
	data, err := s.RoomTypesRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.RoomTypesResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.RoomTypesResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result.ID = data.ID
	result.RoomTypeName = data.RoomTypeName
	result.RoomTypeMaxCapacity = data.RoomTypeMaxCapacity
	result.RoomTypeDesc = data.RoomTypeDesc

	return &result, nil
}

func (s *service) Store(ctx context.Context, payload *dto.CreateRoomTypesRequestBody) (*dto.RoomTypesResponse, error) {
	var result dto.RoomTypesResponse
	isExist, err := s.RoomTypesRepository.ExistByName(ctx, *payload.RoomTypeName)
	if err != nil {
		return &result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	if isExist {
		return &result, res.ErrorBuilder(&res.ErrorConstant.Duplicate, errors.New("room type already exists"))
	}

	data, err := s.RoomTypesRepository.Save(ctx, payload)
	if err != nil {
		return &result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result.ID = data.ID
	result.RoomTypeName = data.RoomTypeName
	result.RoomTypeMaxCapacity = data.RoomTypeMaxCapacity
	result.RoomTypeDesc = data.RoomTypeDesc

	return &result, nil
}

func (s *service) UpdateById(ctx context.Context, payload *dto.UpdateRoomTypesRequestBody) (*dto.RoomTypesResponse, error) {
	roomType, err := s.RoomTypesRepository.FindByID(ctx, *payload.ID)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.RoomTypesResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.RoomTypesResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	_, err = s.RoomTypesRepository.Edit(ctx, &roomType, payload)
	if err != nil {
		return &dto.RoomTypesResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	var result dto.RoomTypesResponse
	result.ID = roomType.ID
	result.RoomTypeName = roomType.RoomTypeName
	result.RoomTypeMaxCapacity = roomType.RoomTypeMaxCapacity
	result.RoomTypeDesc = roomType.RoomTypeDesc

	return &result, nil
}
func (s *service) DeleteById(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.RoomTypesWithCUDResponse, error) {
	roomType, err := s.RoomTypesRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.RoomTypesWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.RoomTypesWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	_, err = s.RoomTypesRepository.Destroy(ctx, &roomType)
	if err != nil {
		return &dto.RoomTypesWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := &dto.RoomTypesWithCUDResponse{
		RoomTypesResponse: dto.RoomTypesResponse{
			ID:                  roomType.ID,
			RoomTypeName:        roomType.RoomTypeName,
			RoomTypeMaxCapacity: roomType.RoomTypeMaxCapacity,
			RoomTypeDesc:        roomType.RoomTypeDesc,
		},
		CreatedAt: roomType.CreatedAt,
		UpdatedAt: roomType.UpdatedAt,
		DeletedAt: roomType.DeletedAt,
	}

	return result, nil
}
