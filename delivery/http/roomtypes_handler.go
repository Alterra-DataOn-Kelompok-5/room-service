package http

import (
	"net/http"
	"strconv"

	"github.com/Alterra-DataOn-Kelompok-5/room-service/domain"
	"github.com/labstack/echo/v4"
)

type RoomTypesHandler struct {
	RoomTypesUsecase domain.RoomTypesUsecase
}

func NewRoomTypesHandler(e *echo.Echo, rtu domain.RoomTypesUsecase) {
	handler := &RoomTypesHandler{
		RoomTypesUsecase: rtu,
	}

	e.GET("/roomtypes", handler.FetchAllRoomTypes)
	e.GET("/roomtypes/:id", handler.FetchRoomTypeByID)
	e.POST("/roomtypes", handler.CreateRoomType)
	e.PUT("/roomtypes/:id", handler.UpdateRoomType)
	e.DELETE("/roomtypes/:id", handler.DeleteRoomType)
}

func (rth *RoomTypesHandler) FetchAllRoomTypes(c echo.Context) error {
	roomTypes, _ := rth.RoomTypesUsecase.FetchAll(c.Request().Context())
	return c.JSON(http.StatusOK, roomTypes)
}

func (rth *RoomTypesHandler) FetchRoomTypeByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	roomTypes, _ := rth.RoomTypesUsecase.FetchByID(c.Request().Context(), id)
	return c.JSON(http.StatusOK, roomTypes)
}

func (rth *RoomTypesHandler) CreateRoomType(c echo.Context) error {
	roomType := domain.RoomTypes{}
	c.Bind(&roomType)

	err := rth.RoomTypesUsecase.Store(c.Request().Context(), &roomType)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Create New Room Type",
	})
}

func (rth *RoomTypesHandler) UpdateRoomType(c echo.Context) error {
	roomType := domain.RoomTypes{}
	c.Bind(&roomType)
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := rth.RoomTypesUsecase.FetchByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Data not found",
		})
	}

	err = rth.RoomTypesUsecase.Update(c.Request().Context(), &roomType, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Update Room Type",
	})
}

func (rth *RoomTypesHandler) DeleteRoomType(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := rth.RoomTypesUsecase.FetchByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Data not found",
		})
	}

	err = rth.RoomTypesUsecase.Delete(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Delete Room Type",
	})
}