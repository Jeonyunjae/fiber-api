package models

type PositionAddressInfo struct {
	UserCode      string  `json:"userCode" gorm:"primaryKey"`
	LocLatitude   float64 `json:"locLatitude"`
	LocLongtitude float64 `json:"locLongtitude"`
}
