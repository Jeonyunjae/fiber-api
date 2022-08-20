package share

import (
	"github.com/jeonyunjae/fiber-api/models"
)

func PositionAddressInfoMapToSlice(rows map[string]models.PositionAddressInfo) []models.PositionAddressInfo {

	positionAddressInfos := []models.PositionAddressInfo{}

	for _, row := range rows {
		var PositionAddressInfo models.PositionAddressInfo
		PositionAddressInfo.UserCode = row.UserCode
		PositionAddressInfo.LocLongtitude = row.LocLongtitude
		PositionAddressInfo.LocLatitude = row.LocLatitude

		positionAddressInfos = append(positionAddressInfos, PositionAddressInfo)

	}
	return positionAddressInfos
}
