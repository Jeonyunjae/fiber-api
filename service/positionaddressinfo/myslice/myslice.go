package myslice

import (
	"github.com/jeonyunjae/fiber-api/models"
	"github.com/jeonyunjae/fiber-api/service/util"
	"github.com/jeonyunjae/fiber-api/util/log"
)

var PositionAddressInfo ULSlice

type ULSlice struct {
	PositionAddressInfoSlice []models.Positionaddressinfo
}

func (ULL *ULSlice) PositionAddressInfoInit(rows []models.Positionaddressinfo) error {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	if rows == nil {
		return log.MyError("Error_PositionAddressInfoInit")
	}
	ULL.PositionAddressInfoSlice = rows

	return nil
}

func (ULL *ULSlice) PositionAddressInfoInsert(ul models.Positionaddressinfo) error {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	ULL.PositionAddressInfoSlice = append(ULL.PositionAddressInfoSlice, ul)

	data, err := ULL.PositionAddressInfoRead(ul)
	if len(data) < 1 || err != nil {
		return err
	}

	return nil
}

func (ULL *ULSlice) PositionAddressInfoRead(ul models.Positionaddressinfo) ([]models.Positionaddressinfo, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	var rows []models.Positionaddressinfo
	var row models.Positionaddressinfo

	for _, row = range ULL.PositionAddressInfoSlice {
		if row.Usercode == ul.Usercode {
			row.Loclatitude = ul.Loclatitude
			row.Loclongtitude = ul.Loclongtitude
			rows = append(rows, row)
			return rows, nil
		}
	}
	return rows, log.MyError("Error_PositionAddressInfoRead")
}

func (ULL *ULSlice) PositionAddressInfoUpdate(ul models.Positionaddressinfo) models.Positionaddressinfo {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	var row models.Positionaddressinfo
	for _, row = range ULL.PositionAddressInfoSlice {
		if row.Usercode == ul.Usercode {
			row.Loclatitude = ul.Loclatitude
			row.Loclongtitude = ul.Loclongtitude
			return row
		}
	}
	return row
}

func (ULL *ULSlice) PositionAddressInfoDelete(ul models.Positionaddressinfo) ULSlice {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	ULL.PositionAddressInfoSlice = util.PositionAddressInfoCsvToStruct()
	return *ULL
}
