package locations

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
	RoomLocationsRepository repository.RoomLocations
}

type Service interface {
	Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto.RoomLocationsResponse], error)
	FindByID(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.RoomLocationsResponse, error)
	Store(ctx context.Context, payload *dto.CreateRoomLocationsRequestBody) (*dto.RoomLocationsResponse, error)
	UpdateById(ctx context.Context, payload *dto.UpdateRoomLocationsRequestBody) (*dto.RoomLocationsResponse, error)
	DeleteById(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.RoomLocationsWithCUDResponse, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		RoomLocationsRepository: f.RoomLocationsRepository,
	}
}

func (s *service) Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto.RoomLocationsResponse], error) {
	roomLocations, info, err := s.RoomLocationsRepository.FindAll(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	var data []dto.RoomLocationsResponse

	for _, roomLoc := range roomLocations {
		data = append(data, dto.RoomLocationsResponse{
			ID:               roomLoc.ID,
			RoomLocationName: roomLoc.RoomLocationName,
			RoomLocationDesc: roomLoc.RoomLocationDesc,
		})

	}

	result := new(pkgdto.SearchGetResponse[dto.RoomLocationsResponse])
	result.Data = data
	result.PaginationInfo = *info

	return result, nil
}
func (s *service) FindByID(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.RoomLocationsResponse, error) {
	var result dto.RoomLocationsResponse
	data, err := s.RoomLocationsRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.RoomLocationsResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.RoomLocationsResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result.ID = data.ID
	result.RoomLocationName = data.RoomLocationName
	result.RoomLocationDesc = data.RoomLocationDesc

	return &result, nil
}

func (s *service) Store(ctx context.Context, payload *dto.CreateRoomLocationsRequestBody) (*dto.RoomLocationsResponse, error) {
	var result dto.RoomLocationsResponse
	isExist, err := s.RoomLocationsRepository.ExistByName(ctx, *payload.RoomLocationName)
	if err != nil {
		return &result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	if isExist {
		return &result, res.ErrorBuilder(&res.ErrorConstant.Duplicate, errors.New("room location already exists"))
	}

	data, err := s.RoomLocationsRepository.Save(ctx, payload)
	if err != nil {
		return &result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result.ID = data.ID
	result.RoomLocationName = data.RoomLocationName
	result.RoomLocationDesc = data.RoomLocationDesc

	return &result, nil
}

func (s *service) UpdateById(ctx context.Context, payload *dto.UpdateRoomLocationsRequestBody) (*dto.RoomLocationsResponse, error) {
	roomLocation, err := s.RoomLocationsRepository.FindByID(ctx, *payload.ID)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.RoomLocationsResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.RoomLocationsResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	_, err = s.RoomLocationsRepository.Edit(ctx, &roomLocation, payload)
	if err != nil {
		return &dto.RoomLocationsResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	var result dto.RoomLocationsResponse
	result.ID = roomLocation.ID
	result.RoomLocationName = roomLocation.RoomLocationName
	result.RoomLocationDesc = roomLocation.RoomLocationDesc

	return &result, nil
}
func (s *service) DeleteById(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.RoomLocationsWithCUDResponse, error) {
	roomLocation, err := s.RoomLocationsRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.RoomLocationsWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.RoomLocationsWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	_, err = s.RoomLocationsRepository.Destroy(ctx, &roomLocation)
	if err != nil {
		return &dto.RoomLocationsWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := &dto.RoomLocationsWithCUDResponse{
		RoomLocationsResponse: dto.RoomLocationsResponse{
			ID:               roomLocation.ID,
			RoomLocationName: roomLocation.RoomLocationName,
			RoomLocationDesc: roomLocation.RoomLocationDesc,
		},
		CreatedAt: roomLocation.CreatedAt,
		UpdatedAt: roomLocation.UpdatedAt,
		DeletedAt: roomLocation.DeletedAt,
	}

	return result, nil
}
