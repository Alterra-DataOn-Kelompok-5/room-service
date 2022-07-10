package model

type Rooms struct {
	Common
	RoomName       string        `json:"room_name"`
	RoomDesc       string        `json:"room_desc"`
	RoomTypeID     int           `json:"room_type_id"`
	RoomTypes      RoomTypes     `json:"room_types" gorm:"foreignKey:RoomTypeID;references:ID"`
	RoomLocationID int           `json:"room_location_id"`
	RoomLocations  RoomLocations `json:"room_locations" gorm:"foreignKey:RoomLocationID;references:ID"`
}
