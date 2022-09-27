package models

type Positionaddressinfo struct {
	Usercode      string  `json:"userCode" gorm:"primaryKey"`
	Loclatitude   float64 `json:"locLatitude"`
	Loclongtitude float64 `json:"locLongtitude"`
	Name          string  `json:"name"`
	Address       string  `json:"address"`
}

type PositionaddressDistanceInfo struct {
	Usercode      string
	Loclatitude   float64
	Loclongtitude float64
	Name          string
	Address       string
	Distance      float64
	Count         int
}
