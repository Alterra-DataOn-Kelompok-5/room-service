package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/Alterra-DataOn-Kelompok-5/room-service/database"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/database/seeder"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/internal/dto"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/internal/factory"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/internal/mocks"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/internal/pkg/enum"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/internal/pkg/util"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	adminClaims             = util.CreateJWTClaims(testEmail, testEmployeeID, testAdminRoleID, testDivisionID)
	db                      = database.GetConnection()
	echoMock                = mocks.EchoMock{E: echo.New()}
	f                       = factory.Factory{RoomTypesRepository: repository.NewRoomTypesRepository(db)}
	roomTypeHandler         = NewHandler(&f)
	testAdminRoleID         = uint(enum.Admin)
	testCreatePayload       = dto.CreateRoomTypesRequestBody{RoomTypeName: &testRoomTypeName, RoomTypeMaxCapacity: &testRoomTypeMaxCapacity, RoomTypeDesc: &testRoomTypeDesc}
	testDivisionID          = uint(enum.Finance)
	testEmail               = "vincentlhubbard@superrito.com"
	testEmployeeID          = uint(1)
	testRoomTypeDesc        = "Ruang meeting kecil dengan kapasitas 1-5 orang"
	testRoomTypeID          = uint(1)
	testRoomTypeMaxCapacity = 5
	testRoomTypeName        = "Small Meeting Room"
	testUpdatePayload       = dto.UpdateRoomTypesRequestBody{ID: &testRoomTypeID, RoomTypeName: &testUpdateRoomTypeName}
	testUpdateRoomTypeName  = "Ruang Meeting Kecil"
	testUserRoleID          = uint(enum.User)
	userClaims              = util.CreateJWTClaims(testEmail, testEmployeeID, testUserRoleID, testDivisionID)
)

func TestRoomTypeHandlerGetInvalidPayload(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/types")
	c.QueryParams().Add("page", "a")
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(roomTypeHandler.Get(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Bad Request")
	}
}

func TestRoomTypeHandlerGetUnauthorized(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	c.SetPath("/api/v1/types")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(roomTypeHandler.Get(c)) {
		asserts.Equal(401, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "unauthorized")
	}
}

func TestRoomTypeHandlerGetSuccess(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	c.SetPath("/api/v1/types")
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(roomTypeHandler.Get(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "room_type_name")
		asserts.Contains(body, "room_type_desc")
		asserts.Contains(body, "room_type_max_capacity")
	}
}

func TestRoomTypeHandlerGetByIdInvalidPayload(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	roleID := "a"

	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/types")
	c.SetParamNames("id")
	c.SetParamValues(roleID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(roomTypeHandler.GetById(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Bad Request")
	}
}

func TestRoomTypeHandlerGetByIdNotFound(t *testing.T) {
	seeder.NewSeeder().DeleteAll()

	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	roleID := strconv.Itoa(int(testAdminRoleID))
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/types")
	c.SetParamNames("id")
	c.SetParamValues(roleID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(roomTypeHandler.GetById(c)) {
		asserts.Equal(404, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Data not found")
	}
}

func TestRoomTypeHandlerGetByIdUnauthorized(t *testing.T) {
	// token, err := util.CreateJWTToken(userClaims)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	c.SetPath("/api/v1/types")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(testAdminRoleID)))
	// c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(roomTypeHandler.GetById(c)) {
		asserts.Equal(401, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "unauthorized")
	}
}

func TestRoomTypeHandlerGetByIdSuccess(t *testing.T) {
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	roleID := strconv.Itoa(int(testAdminRoleID))
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/types")
	c.SetParamNames("id")
	c.SetParamValues(roleID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(roomTypeHandler.GetById(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "room_type_name")
		asserts.Contains(body, "room_type_desc")
		asserts.Contains(body, "room_type_max_capacity")
	}
}

func TestRoomTypeHandlerUpdateByIdInvalidPayload(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodPut, "/", nil)
	roleID := "a"
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/types")
	c.SetParamNames("id")
	c.SetParamValues(roleID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(roomTypeHandler.UpdateById(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Bad Request")
	}
}

func TestRoomTypeHandlerUpdateByIdNotFound(t *testing.T) {
	seeder.NewSeeder().DeleteAll()

	payload, err := json.Marshal(testUpdatePayload)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodPut, "/", bytes.NewBuffer(payload))
	roleID := strconv.Itoa(int(testAdminRoleID))
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/roles")
	c.SetParamNames("id")
	c.SetParamValues(roleID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	c.Request().Header.Set("Content-Type", "application/json")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(roomTypeHandler.UpdateById(c)) {
		asserts.Equal(404, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Data not found")
	}
}
func TestRoomTypeHandlerUpdateByIdUnauthorized(t *testing.T) {
	token, err := util.CreateJWTToken(userClaims)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	c.SetPath("/api/v1/types")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(testAdminRoleID)))
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(roomTypeHandler.UpdateById(c)) {
		asserts.Equal(401, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "unauthorized")
	}
}

func TestRoomTypeHandlerUpdateByIdSuccess(t *testing.T) {
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	payload, err := json.Marshal(testUpdatePayload)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodPut, "/", bytes.NewBuffer(payload))
	roleID := strconv.Itoa(int(testAdminRoleID))
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/roles")
	c.SetParamNames("id")
	c.SetParamValues(roleID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	c.Request().Header.Set("Content-Type", "application/json")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(roomTypeHandler.UpdateById(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "room_type_name")
		asserts.Contains(body, "room_type_desc")
		asserts.Contains(body, "room_type_max_capacity")
	}
}

func TestRoomTypeHandlerDeleteByIdInvalidPayload(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)
	roleID := "a"
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/roles")
	c.SetParamNames("id")
	c.SetParamValues(roleID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(roomTypeHandler.DeleteById(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Bad Request")
	}
}

func TestRoomTypeHandlerDeleteByIdNotFound(t *testing.T) {
	seeder.NewSeeder().DeleteAll()

	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)
	roleID := strconv.Itoa(int(testAdminRoleID))
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/roles")
	c.SetParamNames("id")
	c.SetParamValues(roleID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(roomTypeHandler.DeleteById(c)) {
		asserts.Equal(404, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Data not found")
	}
}

func TestRoomTypeHandlerDeleteByIdUnauthorized(t *testing.T) {
	token, err := util.CreateJWTToken(userClaims)
	if err != nil {
		t.Fatal(err)
	}

	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)

	c.SetPath("/api/v1/types")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(testAdminRoleID)))
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(roomTypeHandler.DeleteById(c)) {
		asserts.Equal(401, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "unauthorized")
	}
}

func TestRoomTypeHandlerDeleteByIdSuccess(t *testing.T) {
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)
	roleID := strconv.Itoa(int(testAdminRoleID))
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/roles")
	c.SetParamNames("id")
	c.SetParamValues(roleID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(roomTypeHandler.DeleteById(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "room_type_name")
		asserts.Contains(body, "room_type_desc")
		asserts.Contains(body, "room_type_max_capacity")
		asserts.Contains(body, "deleted_at")
	}
}

func TestRoomTypeHandlerCreateInvalidPayload(t *testing.T) {
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}
	payload, err := json.Marshal(dto.CreateRoomTypesRequestBody{})
	if err != nil {
		t.Fatal(err)
	}

	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))

	c.SetPath("/api/v1/types")
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	c.Request().Header.Set("Content-Type", "application/json")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(roomTypeHandler.Create(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Invalid parameters or payload")
	}
}

func TestRoomTypeHandlerCreateAlreadyExist(t *testing.T) {
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}
	payload, err := json.Marshal(testCreatePayload)
	if err != nil {
		t.Fatal(err)
	}

	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))

	c.SetPath("/api/v1/types")
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	c.Request().Header.Set("Content-Type", "application/json")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(roomTypeHandler.Create(c)) {
		asserts.Equal(409, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "duplicate")
	}
}

func TestRoomTypeHandlerCreateUnauthorized(t *testing.T) {
	token, err := util.CreateJWTToken(userClaims)
	if err != nil {
		t.Fatal(err)
	}
	payload, err := json.Marshal(testCreatePayload)
	if err != nil {
		t.Fatal(err)
	}

	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))

	c.SetPath("/api/v1/types")
	c.Request().Header.Set("Content-Type", "application/json")
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(roomTypeHandler.Create(c)) {
		asserts.Equal(401, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "unauthorized")
	}
}

func TestRoomTypeHandlerCreateSuccess(t *testing.T) {
	seeder.NewSeeder().DeleteAll()

	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}
	payload, err := json.Marshal(testCreatePayload)
	if err != nil {
		t.Fatal(err)
	}

	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))

	c.SetPath("/api/v1/types")
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	c.Request().Header.Set("Content-Type", "application/json")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(roomTypeHandler.Create(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "room_type_name")
		asserts.Contains(body, "room_type_desc")
		asserts.Contains(body, "room_type_max_capacity")
	}
}
