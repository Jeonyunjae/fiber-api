package myslice

import (
	"github.com/jeonyunjae/fiber-api/models"
	"github.com/jeonyunjae/fiber-api/service/util"
	"github.com/jeonyunjae/fiber-api/util/log"
)

var PositionAddressInfo ULSlice

type ULSlice struct {
	PositionAddressInfoSlice []models.PositionAddressInfo
}

func (ULL *ULSlice) PositionAddressInfosInit(rows []models.PositionAddressInfo) error {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	if rows == nil {
		return log.MyError("Error_PositionAddressInfosInit")
	}
	ULL.PositionAddressInfoSlice = rows

	return nil
}

func (ULL *ULSlice) PositionAddressInfoInsert(ul models.PositionAddressInfo) error {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	ULL.PositionAddressInfoSlice = append(ULL.PositionAddressInfoSlice, ul)

	data, err := ULL.PositionAddressInfoRead(ul)
	if data.UserCode == "" || err != nil {
		return err
	}

	return nil
}

func (ULL *ULSlice) PositionAddressInfoRead(ul models.PositionAddressInfo) (models.PositionAddressInfo, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	var row models.PositionAddressInfo

	for _, row = range ULL.PositionAddressInfoSlice {
		if row.UserCode == ul.UserCode {
			row.LocLatitude = ul.LocLatitude
			row.LocLongtitude = ul.LocLongtitude
			return row, nil
		}
	}
	return row, log.MyError("Error_PositionAddressInfoRead")
}

func (ULL *ULSlice) PositionAddressInfoUpdate(ul models.PositionAddressInfo) models.PositionAddressInfo {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	var row models.PositionAddressInfo
	for _, row = range ULL.PositionAddressInfoSlice {
		if row.UserCode == ul.UserCode {
			row.LocLatitude = ul.LocLatitude
			row.LocLongtitude = ul.LocLongtitude
			return row
		}
	}
	return row
}

func (ULL *ULSlice) PositionAddressInfoDelete(ul models.PositionAddressInfo) ULSlice {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	ULL.PositionAddressInfoSlice = util.PositionAddressInfoCsvToStruct()
	return *ULL
}
