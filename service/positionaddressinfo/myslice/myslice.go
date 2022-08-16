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

func (ULL *ULSlice) PositionAddressInfosInit() ULSlice {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	//ULL.PositionAddressInfoSlice = util.PositionAddressInfoCsvToStruct()

	return *ULL
}

func (ULL *ULSlice) PositionAddressInfoInsert(ul models.PositionAddressInfo) ULSlice {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	ULL.PositionAddressInfoSlice = append(ULL.PositionAddressInfoSlice, ul)

	return *ULL
}

func (ULL *ULSlice) PositionAddressInfoRead(ul models.PositionAddressInfo) ULSlice {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	return *ULL
}

func (ULL *ULSlice) PositionAddressInfoUpdate(ul models.PositionAddressInfo) models.PositionAddressInfo {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	var row models.PositionAddressInfo
	for _, row = range ULL.PositionAddressInfoSlice {
		if row.UserCode == ul.UserCode {
			row.LocLatitude = ul.LocLatitude
			row.LocLongtitue = ul.LocLongtitue
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
