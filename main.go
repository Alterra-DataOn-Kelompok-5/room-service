package main

import (
	"log"
	"os"

	"github.com/Alterra-DataOn-Kelompok-5/room-service/database"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/database/migration"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/delivery/http"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/middleware"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/repository"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/usecase"
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

	roomTypeRepo := repository.NewMysqlRoomTypesRepository(database.GetConnection())
	rtu := usecase.NewRoomTypesUsecase(roomTypeRepo)
	http.NewRoomTypesHandler(e, rtu)

	roomRepo := repository.NewMysqlRoomsRepository(database.GetConnection())
	ru := usecase.NewRoomsUsecase(roomRepo)
	http.NewRoomsHandler(e, ru)

	log.Fatal(e.Start(":" + os.Getenv("SERVICE_PORT")))
}
