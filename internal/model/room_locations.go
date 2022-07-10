package model

type RoomLocations struct {
	Common
	RoomLocationName string `json:"room_location_name"`
	RoomLocationDesc string `json:"room_location_desc"`
}
