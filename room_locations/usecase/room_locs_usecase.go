package usecase

import (
	"context"

	"github.com/Alterra-DataOn-Kelompok-5/room-service/domain"
)

type roomLocationsUsecase struct {
	roomLocationsRepo domain.RoomLocationsRepository
}

func NewRoomLocationsUsecase(rl domain.RoomLocationsRepository) domain.RoomLocationsUsecase {
	return &roomLocationsUsecase{
		roomLocationsRepo: rl,
	}
}

func (rlu *roomLocationsUsecase) FetchAll(c context.Context) (res []domain.RoomLocations, err error) {
	res, err = rlu.roomLocationsRepo.FetchAll(c)
	if err != nil {
		return nil, err
	}
	return
}

func (rlu *roomLocationsUsecase) FetchByID(c context.Context, id int) (res domain.RoomLocations, err error) {
	res, err = rlu.roomLocationsRepo.FetchByID(c, id)
	return
}

func (rlu *roomLocationsUsecase) Store(c context.Context, rl *domain.RoomLocations) (err error) {
	err = rlu.roomLocationsRepo.Store(c, rl)
	return
}

func (rlu *roomLocationsUsecase) Update(c context.Context, rl *domain.RoomLocations, id int) (err error) {
	err = rlu.roomLocationsRepo.Update(c, rl, id)
	return
}

func (rlu *roomLocationsUsecase) Delete(c context.Context, id int) (err error) {
	err = rlu.roomLocationsRepo.Delete(c, id)
	return
}
