package http

import (
	"net/http"
	"strconv"

	"github.com/Alterra-DataOn-Kelompok-5/room-service/domain"
	"github.com/labstack/echo/v4"
)

type RoomsHandler struct {
	RoomsUsecase domain.RoomsUsecase
}

func NewRoomsHandler(e *echo.Echo, ru domain.RoomsUsecase) {
	handler := &RoomsHandler{
		RoomsUsecase: ru,
	}

	e.GET("/rooms", handler.FetchAllRooms)
	e.GET("/rooms/:id", handler.FetchRoomByID)
	e.POST("/rooms", handler.CreateRoom)
	e.PUT("/rooms/:id", handler.UpdateRoom)
	e.DELETE("/rooms/:id", handler.DeleteRoom)
}

func (rh *RoomsHandler) FetchAllRooms(c echo.Context) error {
	rooms, _ := rh.RoomsUsecase.FetchAll(c.Request().Context())
	return c.JSON(http.StatusOK, rooms)
}

func (rh *RoomsHandler) FetchRoomByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	rooms, _ := rh.RoomsUsecase.FetchByID(c.Request().Context(), id)
	return c.JSON(http.StatusOK, rooms)
}

func (rh *RoomsHandler) CreateRoom(c echo.Context) error {
	room := domain.Rooms{}
	c.Bind(&room)

	err := rh.RoomsUsecase.Store(c.Request().Context(), &room)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Create New Room",
	})
}

func (rh *RoomsHandler) UpdateRoom(c echo.Context) error {
	room := domain.Rooms{}
	c.Bind(&room)
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := rh.RoomsUsecase.FetchByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Data not found",
		})
	}

	err = rh.RoomsUsecase.Update(c.Request().Context(), &room, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Update Room",
	})
}

func (rh *RoomsHandler) DeleteRoom(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := rh.RoomsUsecase.FetchByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Data not found",
		})
	}

	err = rh.RoomsUsecase.Delete(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Delete Room",
	})
}
