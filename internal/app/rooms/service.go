package rooms

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
	RoomsRepository         repository.Rooms
	RoomTypesRepository     repository.RoomTypes
	RoomLocationsRepository repository.RoomLocations
}

type Service interface {
	Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto.RoomsDetailResponse], error)
	FindByID(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.RoomsDetailResponse, error)
	Store(ctx context.Context, payload *dto.CreateRoomsRequestBody) (*dto.RoomsDetailResponse, error)
	UpdateById(ctx context.Context, payload *dto.UpdateRoomsRequestBody) (*dto.RoomsDetailResponse, error)
	DeleteById(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.RoomsWithCUDResponse, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		RoomsRepository:         f.RoomsRepository,
		RoomTypesRepository:     f.RoomTypesRepository,
		RoomLocationsRepository: f.RoomLocationsRepository,
	}
}

func (s *service) Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto.RoomsDetailResponse], error) {
	rooms, info, err := s.RoomsRepository.FindAll(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	var data []dto.RoomsDetailResponse

	for _, room := range rooms {
		data = append(data, dto.RoomsDetailResponse{
			RoomsResponse: dto.RoomsResponse{
				ID:       room.ID,
				RoomName: room.RoomName,
				RoomDesc: room.RoomDesc,
			},
			RoomTypes: dto.RoomTypesResponse{
				ID:                  room.RoomTypes.ID,
				RoomTypeName:        room.RoomTypes.RoomTypeName,
				RoomTypeMaxCapacity: room.RoomTypes.RoomTypeMaxCapacity,
				RoomTypeDesc:        room.RoomTypes.RoomTypeDesc,
			},
			RoomLocations: dto.RoomLocationsResponse{
				ID:               room.RoomLocations.ID,
				RoomLocationName: room.RoomLocations.RoomLocationName,
				RoomLocationDesc: room.RoomLocations.RoomLocationDesc,
			},
		})

	}

	result := new(pkgdto.SearchGetResponse[dto.RoomsDetailResponse])
	result.Data = data
	result.PaginationInfo = *info

	return result, nil
}
func (s *service) FindByID(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.RoomsDetailResponse, error) {
	data, err := s.RoomsRepository.FindByID(ctx, payload.ID, true)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.RoomsDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.RoomsDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := &dto.RoomsDetailResponse{
		RoomsResponse: dto.RoomsResponse{
			ID:       data.ID,
			RoomName: data.RoomName,
			RoomDesc: data.RoomDesc,
		},
		RoomTypes: dto.RoomTypesResponse{
			ID:                  data.RoomTypes.ID,
			RoomTypeName:        data.RoomTypes.RoomTypeName,
			RoomTypeMaxCapacity: data.RoomTypes.RoomTypeMaxCapacity,
			RoomTypeDesc:        data.RoomTypes.RoomTypeDesc,
		},
		RoomLocations: dto.RoomLocationsResponse{
			ID:               data.RoomLocations.ID,
			RoomLocationName: data.RoomLocations.RoomLocationName,
			RoomLocationDesc: data.RoomLocations.RoomLocationDesc,
		},
	}

	return result, nil
}

func (s *service) Store(ctx context.Context, payload *dto.CreateRoomsRequestBody) (*dto.RoomsDetailResponse, error) {
	isExist, err := s.RoomsRepository.ExistByName(ctx, *payload.RoomName)
	if err != nil {
		return &dto.RoomsDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	if isExist {
		return &dto.RoomsDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.Duplicate, errors.New("room already exists"))
	}

	isRoomTypeActive, err := s.RoomTypesRepository.ExistByID(ctx, *payload.RoomTypeID)
	if err != nil {
		return &dto.RoomsDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	if !isRoomTypeActive {
		return &dto.RoomsDetailResponse{}, res.CustomErrorBuilder(res.ErrorConstant.NotFound.Code, res.E_NOT_FOUND, "Room type not found")
	}

	isRoomLocationActive, err := s.RoomLocationsRepository.ExistByID(ctx, *payload.RoomLocationID)
	if err != nil {
		return &dto.RoomsDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	if !isRoomLocationActive {
		return &dto.RoomsDetailResponse{}, res.CustomErrorBuilder(res.ErrorConstant.NotFound.Code, res.E_NOT_FOUND, "Room location not found")
	}

	data, err := s.RoomsRepository.Save(ctx, payload)
	if err != nil {
		return &dto.RoomsDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	room, err := s.RoomsRepository.FindByID(ctx, data.ID, true)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.RoomsDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.RoomsDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := &dto.RoomsDetailResponse{
		RoomsResponse: dto.RoomsResponse{
			ID:       room.ID,
			RoomName: room.RoomName,
			RoomDesc: room.RoomDesc,
		},
		RoomTypes: dto.RoomTypesResponse{
			ID:                  room.RoomTypes.ID,
			RoomTypeName:        room.RoomTypes.RoomTypeName,
			RoomTypeMaxCapacity: room.RoomTypes.RoomTypeMaxCapacity,
			RoomTypeDesc:        room.RoomTypes.RoomTypeDesc,
		},
		RoomLocations: dto.RoomLocationsResponse{
			ID:               room.RoomLocations.ID,
			RoomLocationName: room.RoomLocations.RoomLocationName,
			RoomLocationDesc: room.RoomLocations.RoomLocationDesc,
		},
	}

	return result, nil
}

func (s *service) UpdateById(ctx context.Context, payload *dto.UpdateRoomsRequestBody) (*dto.RoomsDetailResponse, error) {
	room, err := s.RoomsRepository.FindByID(ctx, *payload.ID, false)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.RoomsDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.RoomsDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	isRoomTypeActive, err := s.RoomTypesRepository.ExistByID(ctx, *payload.RoomTypeID)
	if err != nil {
		return &dto.RoomsDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	if !isRoomTypeActive {
		return &dto.RoomsDetailResponse{}, res.CustomErrorBuilder(res.ErrorConstant.NotFound.Code, res.E_NOT_FOUND, "Room type not found")
	}

	isRoomLocationActive, err := s.RoomLocationsRepository.ExistByID(ctx, *payload.RoomLocationID)
	if err != nil {
		return &dto.RoomsDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	if !isRoomLocationActive {
		return &dto.RoomsDetailResponse{}, res.CustomErrorBuilder(res.ErrorConstant.NotFound.Code, res.E_NOT_FOUND, "Room location not found")
	}

	_, err = s.RoomsRepository.Edit(ctx, &room, payload)
	if err != nil {
		return &dto.RoomsDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := &dto.RoomsDetailResponse{
		RoomsResponse: dto.RoomsResponse{
			ID:       room.ID,
			RoomName: room.RoomName,
			RoomDesc: room.RoomDesc,
		},
		RoomTypes: dto.RoomTypesResponse{
			ID:                  room.RoomTypes.ID,
			RoomTypeName:        room.RoomTypes.RoomTypeName,
			RoomTypeMaxCapacity: room.RoomTypes.RoomTypeMaxCapacity,
			RoomTypeDesc:        room.RoomTypes.RoomTypeDesc,
		},
		RoomLocations: dto.RoomLocationsResponse{
			ID:               room.RoomLocations.ID,
			RoomLocationName: room.RoomLocations.RoomLocationName,
			RoomLocationDesc: room.RoomLocations.RoomLocationDesc,
		},
	}

	return result, nil
}
func (s *service) DeleteById(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.RoomsWithCUDResponse, error) {
	room, err := s.RoomsRepository.FindByID(ctx, payload.ID, true)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.RoomsWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.RoomsWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	_, err = s.RoomsRepository.Destroy(ctx, &room)
	if err != nil {
		return &dto.RoomsWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := &dto.RoomsWithCUDResponse{
		RoomsResponse: dto.RoomsResponse{
			ID:       room.ID,
			RoomName: room.RoomName,
			RoomDesc: room.RoomDesc,
		},
		CreatedAt: room.CreatedAt,
		UpdatedAt: room.UpdatedAt,
		DeletedAt: room.DeletedAt,
	}

	return result, nil
}
