package models

import "time"

type PositionAddressInfo struct {
	UserCode     uint64  `json:"id" gorm:"primaryKey"`
	LocLatitude  float64 `json:"Lon"`
	LocLongtitue float64 `json:"Lat"`
	CreatedAt    time.Time
}
