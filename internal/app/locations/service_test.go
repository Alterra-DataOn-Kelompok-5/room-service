package locations

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
	ctx                 = context.Background()
	locationService     = NewService(factory.NewFactory())
	testFindAllPayload  = pkgdto.SearchGetRequest{}
	testFindByIdPayload = pkgdto.ByIDRequest{ID: 1}
)

func TestLocationServiceFindAllSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	asserts := assert.New(t)
	res, err := locationService.Find(ctx, &testFindAllPayload)
	if err != nil {
		t.Fatal(err)
	}

	asserts.Len(res.Data, 3)
	for _, val := range res.Data {
		asserts.NotEmpty(val.ID)
		asserts.NotEmpty(val.RoomLocationName)
		asserts.NotEmpty(val.RoomLocationDesc)
	}
}
func TestLocationServiceFindByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	asserts := assert.New(t)
	res, err := locationService.FindByID(ctx, &testFindByIdPayload)
	if err != nil {
		t.Fatal(err)
	}

	asserts.Equal(uint(1), res.ID)
}

func TestLocationServiceFindByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	asserts := assert.New(t)
	_, err := locationService.FindByID(ctx, &testFindByIdPayload)
	if err != nil {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestLocationServiceUpdateByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	asserts := assert.New(t)
	res, err := locationService.UpdateById(ctx, &testUpdatePayload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.Equal(testUpdateRoomLocationDesc, res.RoomLocationDesc)
}

func TestLocationServiceUpdateByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	asserts := assert.New(t)
	_, err := locationService.UpdateById(ctx, &testUpdatePayload)
	if err != nil {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestLocationServiceDeleteByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	asserts := assert.New(t)
	res, err := locationService.DeleteById(ctx, &testFindByIdPayload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.NotNil(res.DeletedAt)
}

func TestLocationServiceDeleteByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	asserts := assert.New(t)
	_, err := locationService.DeleteById(ctx, &testFindByIdPayload)
	if asserts.Error(err) {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestLocationServiceCreateLocationSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	asserts := assert.New(t)
	res, err := locationService.Store(ctx, &testCreatePayload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.NotEmpty(res.ID)
	asserts.Equal(*testCreatePayload.RoomLocationName, res.RoomLocationName)
	asserts.Equal(*testCreatePayload.RoomLocationDesc, res.RoomLocationDesc)
}

func TestLocationServiceCreateLocationAlreadyExist(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	asserts := assert.New(t)
	_, err := locationService.Store(ctx, &testCreatePayload)
	if asserts.Error(err) {
		asserts.Equal(err.Error(), "error code 409")
	}
}
