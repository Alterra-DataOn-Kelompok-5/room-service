package usecase

import (
	"context"
	"errors"

	"github.com/Alterra-DataOn-Kelompok-5/room-service/domain"
)

type roomsUsecase struct {
	roomsRepo     domain.RoomsRepository
	roomTypesRepo domain.RoomTypesRepository
	roomLocsRepo  domain.RoomLocationsRepository
}

func NewRoomsUsecase(r domain.RoomsRepository, rt domain.RoomTypesRepository, rl domain.RoomLocationsRepository) domain.RoomsUsecase {
	return &roomsUsecase{
		roomsRepo:     r,
		roomTypesRepo: rt,
		roomLocsRepo:  rl,
	}
}

func (ru *roomsUsecase) FetchAll(c context.Context) (res []domain.Rooms, err error) {
	res, err = ru.roomsRepo.FetchAll(c)
	if err != nil {
		return nil, err
	}
	return
}

func (ru *roomsUsecase) FetchByID(c context.Context, id int) (res domain.Rooms, err error) {
	res, err = ru.roomsRepo.FetchByID(c, id)
	return
}

func (ru *roomsUsecase) Store(c context.Context, r *domain.Rooms) error {

	roomTypeID := r.RoomTypeID
	roomLocID := r.RoomLocationID

	_, err := ru.roomTypesRepo.FetchByID(c, roomTypeID)
	if err != nil {
		err = errors.New("record not found: room_types")
		return err
	}

	_, err = ru.roomLocsRepo.FetchByID(c, roomLocID)
	if err != nil {
		err = errors.New("record not found: room_locations")
		return err
	}

	err = ru.roomsRepo.Store(c, r)
	return err
}

func (ru *roomsUsecase) Update(c context.Context, r *domain.Rooms, id int) error {
	roomTypeID := r.RoomTypeID
	roomLocID := r.RoomLocationID

	_, err := ru.roomTypesRepo.FetchByID(c, roomTypeID)
	if err != nil {
		err = errors.New("record not found: room_types")
		return err
	}

	_, err = ru.roomLocsRepo.FetchByID(c, roomLocID)
	if err != nil {
		err = errors.New("record not found: room_locations")
		return err
	}
	err = ru.roomsRepo.Update(c, r, id)
	return err
}

func (ru *roomsUsecase) Delete(c context.Context, id int) (err error) {
	err = ru.roomsRepo.Delete(c, id)
	return
}
