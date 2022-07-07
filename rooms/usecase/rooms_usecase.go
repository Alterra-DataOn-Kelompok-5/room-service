package usecase

import (
	"context"

	"github.com/Alterra-DataOn-Kelompok-5/room-service/domain"
)

type roomsUsecase struct {
	roomsRepo domain.RoomsRepository
}

func NewRoomsUsecase(r domain.RoomsRepository) domain.RoomsUsecase {
	return &roomsUsecase{
		roomsRepo: r,
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

func (ru *roomsUsecase) Store(c context.Context, r *domain.Rooms) (err error) {
	err = ru.roomsRepo.Store(c, r)
	return
}

func (ru *roomsUsecase) Update(c context.Context, r *domain.Rooms, id int) (err error) {
	err = ru.roomsRepo.Update(c, r, id)
	return
}

func (ru *roomsUsecase) Delete(c context.Context, id int) (err error) {
	err = ru.roomsRepo.Delete(c, id)
	return
}
