package models

type Positionaddressinfo struct {
	Usercode      string  `json:"userCode" gorm:"primaryKey"`
	Loclatitude   float64 `json:"locLatitude"`
	Loclongtitude float64 `json:"locLongtitude"`
}
