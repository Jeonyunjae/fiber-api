package models

import (
	"time"
)

type ULDecimalTreeNode struct {
	PositionAddressInfoInfo           map[uint]PositionAddressInfo
	PositionAddressInfoDecimalLatTree map[uint]ULDecimalTreeNode
}

type PositionAddressInfo struct {
	ID        uint64  `json:"id" gorm:"primaryKey"`
	Lon       float64 `json:"Lon"`
	Lat       float64 `json:"Lat"`
	CreatedAt time.Time
}
