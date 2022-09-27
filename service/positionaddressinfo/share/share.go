package share

import (
	"github.com/jeonyunjae/fiber-api/models"
)

func PositionAddressInfoMapToSlice(rows map[string]models.Positionaddressinfo) []models.Positionaddressinfo {

	positionAddressInfo := []models.Positionaddressinfo{}

	for _, row := range rows {
		var PositionAddressInfo models.Positionaddressinfo
		PositionAddressInfo.Usercode = row.Usercode
		PositionAddressInfo.Loclongtitude = row.Loclongtitude
		PositionAddressInfo.Loclatitude = row.Loclatitude
		PositionAddressInfo.Name = row.Name
		PositionAddressInfo.Address = row.Address

		positionAddressInfo = append(positionAddressInfo, PositionAddressInfo)

	}
	return positionAddressInfo
}
