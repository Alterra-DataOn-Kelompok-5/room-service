package model

type RoomTypes struct {
	Common
	RoomTypeName        string `json:"room_type_name" gorm:"varchar;not_null;unique"`
	RoomTypeMaxCapacity int    `json:"room_type_max_capacity"`
	RoomTypeDesc        string `json:"room_type_desc" gorm:"type:text"`
}
