package rooms

import (
	"context"
	"testing"

	"github.com/Alterra-DataOn-Kelompok-5/room-service/database"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/database/seeder"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/internal/dto"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/internal/factory"
	pkgdto "github.com/Alterra-DataOn-Kelompok-5/room-service/pkg/dto"
	"github.com/stretchr/testify/assert"
)

var (
	ctx         = context.Background()
	roomService = NewService(factory.NewFactory())
)

func TestRoomServiceFindAllSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	var (
		asserts = assert.New(t)
		payload = pkgdto.SearchGetRequest{}
	)

	res, err := roomService.Find(ctx, &payload)
	if err != nil {
		t.Fatal(err)
	}

	asserts.Len(res.Data, 3)
	for _, val := range res.Data {
		asserts.NotEmpty(val.RoomName)
		asserts.NotEmpty(val.ID)
	}
}
func TestRoomServiceFindByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	var (
		asserts = assert.New(t)
		payload = pkgdto.ByIDRequest{ID: 1}
	)

	res, err := roomService.FindByID(ctx, &payload)
	if err != nil {
		t.Fatal(err)
	}

	asserts.Equal(uint(1), res.ID)
}

func TestRoomServiceFindByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	var (
		asserts = assert.New(t)
		payload = pkgdto.ByIDRequest{ID: 1}
	)

	_, err := roomService.FindByID(ctx, &payload)
	if err != nil {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestRoomServiceUpdataByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	var (
		asserts    = assert.New(t)
		id         = uint(1)
		name       = "Tulip"
		desc       = "test room tulip"
		roomTypeID = uint(2)
		roomLocID  = uint(2)
		payload    = dto.UpdateRoomsRequestBody{
			ID:             &id,
			RoomName:       &name,
			RoomDesc:       &desc,
			RoomTypeID:     &roomTypeID,
			RoomLocationID: &roomLocID,
		}
	)
	res, err := roomService.UpdateById(ctx, &payload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.Equal(name, res.RoomName)
}

func TestRoomServiceUpdateByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	var (
		asserts = assert.New(t)
		id      = uint(1)
		name    = "Tulips"
		payload = dto.UpdateRoomsRequestBody{
			ID:       &id,
			RoomName: &name,
		}
	)

	_, err := roomService.UpdateById(ctx, &payload)
	if err != nil {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestRoomServiceDeleteByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	var (
		asserts = assert.New(t)
		id      = uint(1)
		payload = pkgdto.ByIDRequest{ID: id}
	)

	res, err := roomService.DeleteById(ctx, &payload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.NotNil(res.DeletedAt)
}

func TestRoomServiceDeleteByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	var (
		asserts = assert.New(t)
		id      = uint(10)
		payload = pkgdto.ByIDRequest{ID: id}
	)

	_, err := roomService.DeleteById(ctx, &payload)
	if asserts.Error(err) {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestRoomServiceCreateRoomSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	var (
		asserts = assert.New(t)
		name    = "Mawar"
		payload = dto.CreateRoomsRequestBody{
			RoomName:       &name,
			RoomDesc:       &testRoomDesc,
			RoomTypeID:     &testRoomTypeID,
			RoomLocationID: &testRoomLocationID,
		}
	)

	res, err := roomService.Store(ctx, &payload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.NotEmpty(res.ID)
	asserts.Equal(*payload.RoomName, res.RoomName)
}

func TestRoomServiceCreateRoomAlreadyExist(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	var (
		asserts = assert.New(t)
		name    = "Melati"
		payload = dto.CreateRoomsRequestBody{
			RoomName: &name,
		}
	)

	_, err := roomService.Store(ctx, &payload)
	if asserts.Error(err) {
		asserts.Equal(err.Error(), "error code 409")
	}
}
