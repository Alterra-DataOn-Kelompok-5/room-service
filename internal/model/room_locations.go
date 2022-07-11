package model

type RoomLocations struct {
	Common
	RoomLocationName string `json:"room_location_name" gorm:"varchar;not_null;unique"`
	RoomLocationDesc string `json:"room_location_desc"`
}
