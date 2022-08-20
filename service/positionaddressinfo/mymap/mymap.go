package mymap

import (
	"github.com/jeonyunjae/fiber-api/models"
	"github.com/jeonyunjae/fiber-api/util/log"
)

var PositionAddressInfo ULMap

type ULMap struct {
	PositionAddressInfoMap map[string]models.PositionAddressInfo
}

func (ULM *ULMap) PositionAddressInfosInit(rows map[string]models.PositionAddressInfo) error {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	if rows == nil {
		return log.MyError("Error_PositionAddressInfosInit")
	}
	ULM.PositionAddressInfoMap = rows

	return nil
}

func (ULM *ULMap) PositionAddressInfoInsert(ul models.PositionAddressInfo) error {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	ULM.PositionAddressInfoMap[ul.UserCode] = ul

	data, err := ULM.PositionAddressInfoRead(ul)
	if data.UserCode == "" || err != nil {
		return err
	}

	return nil
}

func (ULM *ULMap) PositionAddressInfoRead(ul models.PositionAddressInfo) (models.PositionAddressInfo, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	row := ULM.PositionAddressInfoMap[ul.UserCode]
	if row.UserCode == "" {
		return row, log.MyError("NotFound")
	}
	return row, nil
}

func (ULM *ULMap) PositionAddressInfoUpdate(ul models.PositionAddressInfo) (bool, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	var row models.PositionAddressInfo
	for _, row = range ULM.PositionAddressInfoMap {
		if row.UserCode == ul.UserCode {
			row.LocLatitude = ul.LocLatitude
			row.LocLongtitude = ul.LocLongtitude
			return true, nil
		}
	}
	return false, log.MyError("NotFound")
}

func (ULM *ULMap) PositionAddressInfoDelete(ul models.PositionAddressInfo) (bool, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	return false, nil
}
