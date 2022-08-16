package mymap

import (
	"github.com/jeonyunjae/fiber-api/models"
	"github.com/jeonyunjae/fiber-api/service/util"
	"github.com/jeonyunjae/fiber-api/util/log"
)

var PositionAddressInfo ULMap

type ULMap struct {
	PositionAddressInfoMap map[int]models.PositionAddressInfo
}

func (ULM *ULMap) PositionAddressInfosInit() ULMap {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	ULM.PositionAddressInfoMap = util.PositionAddressInfoCsvToMap()

	return *ULM
}

func (ULM *ULMap) PositionAddressInfoInsert(ul models.PositionAddressInfo) ULMap {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	ULM.PositionAddressInfoMap[int(ul.UserCode)] = ul

	return *ULM
}

func (ULM *ULMap) PositionAddressInfoRead(ul models.PositionAddressInfo) ULMap {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	return *ULM
}

func (ULM *ULMap) PositionAddressInfoUpdate(ul models.PositionAddressInfo) (bool, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	var row models.PositionAddressInfo
	for _, row = range ULM.PositionAddressInfoMap {
		if row.UserCode == ul.UserCode {
			row.LocLatitude = ul.LocLatitude
			row.LocLongtitue = ul.LocLongtitue
			return true, nil
		}
	}
	return false, log.MyError("NotFound")
}

func (ULM *ULMap) PositionAddressInfoDelete(ul models.PositionAddressInfo) (bool, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	return false, nil
}
