package models

import "time"

type ULList struct {
	UserlocationList []UserLocation
}

type UserLocation struct {
	ID        uint64  `json:"id" gorm:"primaryKey"`
	Lon       float64 `json:"Lon"`
	Lat       float64 `json:"Lat"`
	CityCode  uint64  `json:"CityCode"`
	CreatedAt time.Time
}
