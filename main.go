package main

import (
	"log"
	"os"

	"github.com/Alterra-DataOn-Kelompok-5/room-service/database"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/database/migration"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/middleware"

	_roomTypesHttp "github.com/Alterra-DataOn-Kelompok-5/room-service/room_types/delivery/http"
	_roomTypesRepo "github.com/Alterra-DataOn-Kelompok-5/room-service/room_types/repository"
	_roomTypesUc "github.com/Alterra-DataOn-Kelompok-5/room-service/room_types/usecase"

	_roomLocationsHttp "github.com/Alterra-DataOn-Kelompok-5/room-service/room_locations/delivery/http"
	_roomLocationsRepo "github.com/Alterra-DataOn-Kelompok-5/room-service/room_locations/repository"
	_roomLocationsUc "github.com/Alterra-DataOn-Kelompok-5/room-service/room_locations/usecase"

	_roomsHttp "github.com/Alterra-DataOn-Kelompok-5/room-service/rooms/delivery/http"
	_roomsRepo "github.com/Alterra-DataOn-Kelompok-5/room-service/rooms/repository"
	_roomsUc "github.com/Alterra-DataOn-Kelompok-5/room-service/rooms/usecase"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	errGoEnv := godotenv.Load()
	if errGoEnv != nil {
		// log.Fatal("Error loading .env file")
		panic(errGoEnv)
	}
}

func main() {
	database.CreateConnection()
	migration.Migrate()

	e := echo.New()

	middleware.Init(e)

	roomTypeRepo := _roomTypesRepo.NewMysqlRoomTypesRepository(database.GetConnection())
	rtu := _roomTypesUc.NewRoomTypesUsecase(roomTypeRepo)
	_roomTypesHttp.NewRoomTypesHandler(e, rtu)

	roomLocationRepo := _roomLocationsRepo.NewMysqlRoomLocationsRepository(database.GetConnection())
	rlu := _roomLocationsUc.NewRoomLocationsUsecase(roomLocationRepo)
	_roomLocationsHttp.NewRoomsHandler(e, rlu)

	roomRepo := _roomsRepo.NewMysqlRoomsRepository(database.GetConnection())
	ru := _roomsUc.NewRoomsUsecase(roomRepo)
	_roomsHttp.NewRoomsHandler(e, ru)

	log.Fatal(e.Start(":" + os.Getenv("SERVICE_PORT")))
}
