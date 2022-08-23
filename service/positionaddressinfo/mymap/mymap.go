package mymap

import (
	"github.com/jeonyunjae/fiber-api/models"
	"github.com/jeonyunjae/fiber-api/util/log"
)

var PositionAddressInfo ULMap

type ULMap struct {
	PositionAddressInfoMap map[string]models.Positionaddressinfo
}

func (ULM *ULMap) PositionAddressInfoInit(rows map[string]models.Positionaddressinfo) error {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	if rows == nil {
		return log.MyError("Error_PositionAddressInfoInit")
	}
	ULM.PositionAddressInfoMap = rows

	return nil
}

func (ULM *ULMap) PositionAddressInfoInsert(ul models.Positionaddressinfo) error {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	ULM.PositionAddressInfoMap[ul.Usercode] = ul

	data, err := ULM.PositionAddressInfoRead(ul)
	if len(data) < 1 || err != nil {
		return err
	}

	return nil
}

func (ULM *ULMap) PositionAddressInfoRead(ul models.Positionaddressinfo) (map[string]models.Positionaddressinfo, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	rows := make(map[string]models.Positionaddressinfo)

	row := ULM.PositionAddressInfoMap[ul.Usercode]
	if row.Usercode != "" {
		rows[row.Usercode] = row
		return rows, nil
	}
	return rows, log.MyError("NotFound")
}

func (ULM *ULMap) PositionAddressInfoUpdate(ul models.Positionaddressinfo) (bool, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	// var row models.Positionaddressinfo
	// for _, row = range ULM.PositionAddressInfoMap {
	// 	if row.Usercode == ul.Usercode {
	// 		row.Loclatitude = ul.Loclatitude
	// 		row.Loclongtitude = ul.Loclongtitude
	// 		return true, nil
	// 	}
	// }

	ULM.PositionAddressInfoMap[ul.Usercode] = ul

	return true, nil
}

func (ULM *ULMap) PositionAddressInfoDelete(ul models.Positionaddressinfo) (bool, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	delete(ULM.PositionAddressInfoMap, ul.Usercode)

	return true, nil
}
