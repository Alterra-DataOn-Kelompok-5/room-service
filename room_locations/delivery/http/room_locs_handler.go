package http

import (
	"net/http"
	"strconv"

	"github.com/Alterra-DataOn-Kelompok-5/room-service/domain"
	"github.com/labstack/echo/v4"
)

type RoomLocationsHandler struct {
	RoomLocationsUsecase domain.RoomLocationsUsecase
}

func NewRoomsHandler(e *echo.Echo, ru domain.RoomLocationsUsecase) {
	handler := &RoomLocationsHandler{
		RoomLocationsUsecase: ru,
	}

	e.GET("/locations", handler.FetchAllRoomlocations)
	e.GET("/locations/:id", handler.FetchRoomLocationByID)
	e.POST("/locations", handler.CreateRoomLocation)
	e.PUT("/locations/:id", handler.UpdateRoomLocation)
	e.DELETE("/locations/:id", handler.DeleteRoomLocation)
}

func (rlh *RoomLocationsHandler) FetchAllRoomlocations(c echo.Context) error {
	roomLocations, _ := rlh.RoomLocationsUsecase.FetchAll(c.Request().Context())
	return c.JSON(http.StatusOK, roomLocations)
}

func (rlh *RoomLocationsHandler) FetchRoomLocationByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	roomLocations, err := rlh.RoomLocationsUsecase.FetchByID(c.Request().Context(), id)
	if err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, roomLocations)
}

func (rlh *RoomLocationsHandler) CreateRoomLocation(c echo.Context) error {
	roomLocation := domain.RoomLocations{}
	c.Bind(&roomLocation)

	err := rlh.RoomLocationsUsecase.Store(c.Request().Context(), &roomLocation)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Create New Room Location",
	})
}

func (rlh *RoomLocationsHandler) UpdateRoomLocation(c echo.Context) error {
	roomLocation := domain.RoomLocations{}
	c.Bind(&roomLocation)
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := rlh.RoomLocationsUsecase.FetchByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Data not found",
		})
	}

	err = rlh.RoomLocationsUsecase.Update(c.Request().Context(), &roomLocation, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Update Room Location",
	})
}

func (rlh *RoomLocationsHandler) DeleteRoomLocation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := rlh.RoomLocationsUsecase.FetchByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Data not found",
		})
	}

	err = rlh.RoomLocationsUsecase.Delete(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Delete Room Location",
	})
}
