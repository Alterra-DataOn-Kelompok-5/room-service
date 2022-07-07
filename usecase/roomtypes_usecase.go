package usecase

import (
	"context"

	"github.com/Alterra-DataOn-Kelompok-5/room-service/domain"
)

type roomTypesUsecase struct {
	roomTypesRepo domain.RoomTypesRepository
}

func NewRoomTypesUsecase(rt domain.RoomTypesRepository) domain.RoomTypesUsecase {
	return &roomTypesUsecase{
		roomTypesRepo: rt,
	}
}

func (rtu *roomTypesUsecase) FetchAll(c context.Context) (res []domain.RoomTypes, err error) {
	res, err = rtu.roomTypesRepo.FetchAll(c)
	if err != nil {
		return nil, err
	}
	return
}

func (rtu *roomTypesUsecase) FetchByID(c context.Context, id int) (res domain.RoomTypes, err error) {
	res, err = rtu.roomTypesRepo.FetchByID(c, id)
	return
}

func (rtu *roomTypesUsecase) Store(c context.Context, rt *domain.RoomTypes) (err error) {
	err = rtu.roomTypesRepo.Store(c, rt)
	return
}

func (rtu *roomTypesUsecase) Update(c context.Context, rt *domain.RoomTypes, id int) (err error) {
	err = rtu.roomTypesRepo.Update(c, rt, id)
	return
}

func (rtu *roomTypesUsecase) Delete(c context.Context, id int) (err error) {
	err = rtu.roomTypesRepo.Delete(c, id)
	return
}
