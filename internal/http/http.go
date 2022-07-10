package http

import (
	"github.com/Alterra-DataOn-Kelompok-5/room-service/internal/app/locations"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/internal/app/rooms"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/internal/app/types"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/internal/factory"
	"github.com/Alterra-DataOn-Kelompok-5/room-service/pkg/util"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	e.GET("/status", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "OK"})
	})
	v1 := e.Group("/api/v1")
	rooms.NewHandler(f).Route(v1.Group("/rooms"))
	types.NewHandler(f).Route(v1.Group("/types"))
	locations.NewHandler(f).Route(v1.Group("/locations"))
}
