package types

import (
	"context"
	"testing"

	"github.com/Alterra-DataOn-Kelompok-5/room-service/database"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/database/seeder"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/internal/factory"
	pkgdto "github.com/Alterra-DataOn-Kelompok-5/room-service/pkg/dto"
	"github.com/stretchr/testify/assert"
)

var (
	ctx         = context.Background()
	roomTypeService = NewService(factory.NewFactory())
	testFindAllPayload  = pkgdto.SearchGetRequest{}
	testFindByIdPayload = pkgdto.ByIDRequest{ID: 1}
)

func TestRoomTypeServiceFindAllSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	asserts := assert.New(t)
	res, err := roomTypeService.Find(ctx, &testFindAllPayload)
	if err != nil {
		t.Fatal(err)
	}

	asserts.Len(res.Data, 2)
	for _, val := range res.Data {
		asserts.NotEmpty(val.RoomTypeName)
		asserts.NotEmpty(val.RoomTypeDesc)
		asserts.NotEmpty(val.RoomTypeMaxCapacity)
		asserts.NotEmpty(val.ID)
	}
}
func TestRoomTypeServiceFindByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	asserts := assert.New(t)
	res, err := roomTypeService.FindByID(ctx, &testFindByIdPayload)
	if err != nil {
		t.Fatal(err)
	}

	asserts.Equal(uint(1), res.ID)
}

func TestRoomTypeServiceFindByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	asserts := assert.New(t)
	_, err := roomTypeService.FindByID(ctx, &testFindByIdPayload)
	if err != nil {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestRoomTypeServiceUpdateByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	asserts := assert.New(t)
	res, err := roomTypeService.UpdateById(ctx, &testUpdatePayload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.Equal(testUpdateRoomTypeName, res.RoomTypeName)
}

func TestRoomTypeServiceUpdateByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	
	asserts := assert.New(t)
	_, err := roomTypeService.UpdateById(ctx, &testUpdatePayload)
	if err != nil {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestRoomTypeServiceDeleteByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	asserts := assert.New(t)
	res, err := roomTypeService.DeleteById(ctx, &testFindByIdPayload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.NotNil(res.DeletedAt)
}

func TestRoomTypeServiceDeleteByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	asserts := assert.New(t)

	_, err := roomTypeService.DeleteById(ctx, &testFindByIdPayload)
	if asserts.Error(err) {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestRoomTypeServiceCreateRoomTypeSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	asserts := assert.New(t)
	res, err := roomTypeService.Store(ctx, &testCreatePayload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.NotEmpty(res.ID)
	asserts.Equal(*testCreatePayload.RoomTypeName, res.RoomTypeName)
	asserts.Equal(*testCreatePayload.RoomTypeDesc, res.RoomTypeDesc)
	asserts.Equal(*testCreatePayload.RoomTypeMaxCapacity, res.RoomTypeMaxCapacity)
}

func TestRoomTypeServiceCreateRoomTypeAlreadyExist(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	
	asserts := assert.New(t)
	_, err := roomTypeService.Store(ctx, &testCreatePayload)
	if asserts.Error(err) {
		asserts.Equal(err.Error(), "error code 409")
	}
}
