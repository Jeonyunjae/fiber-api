package mymap

import (
	"sort"

	"github.com/jeonyunjae/fiber-api/models"
	"github.com/jeonyunjae/fiber-api/service/util"
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

func (ULM *ULMap) PositionAddressInfoReads(ul models.PositionaddressDistanceInfo) ([]models.PositionaddressDistanceInfo, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	var sortData []models.PositionaddressDistanceInfo

	for _, row := range ULM.PositionAddressInfoMap {
		var data models.PositionaddressDistanceInfo
		distance := util.GetDistance(ul.Loclongtitude, row.Loclongtitude, ul.Loclatitude, row.Loclatitude)

		data.Usercode = row.Usercode
		data.Loclongtitude = row.Loclongtitude
		data.Loclatitude = row.Loclatitude
		data.Address = row.Address
		data.Name = row.Name
		data.Distance = distance
		sortData = append(sortData, data)
	}
	sort.Slice(sortData, func(i, j int) bool {
		return sortData[i].Distance < sortData[j].Distance
	})
	tempCount := ul.Count
	if len(sortData) < tempCount {
		tempCount = len(sortData)
	}
	sortData = sortData[:tempCount]
	return sortData, nil
}

func (ULM *ULMap) PositionAddressInfoReadsRange(ul models.PositionaddressDistanceInfo) ([]models.PositionaddressDistanceInfo, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	var sortData []models.PositionaddressDistanceInfo

	for _, row := range ULM.PositionAddressInfoMap {
		distance := util.GetDistance(ul.Loclongtitude, row.Loclongtitude, ul.Loclatitude, row.Loclatitude)
		sortData = append(sortData, models.PositionaddressDistanceInfo{Usercode: row.Usercode, Loclatitude: row.Loclatitude, Loclongtitude: row.Loclongtitude, Name: row.Name, Address: row.Address, Distance: distance})
	}
	sort.Slice(sortData, func(i, j int) bool {
		return sortData[i].Distance < sortData[j].Distance
	})
	tempCount := ul.Count
	if len(sortData) < tempCount {
		tempCount = len(sortData)
	}
	sortData = sortData[:tempCount]
	return sortData, nil
}

func (ULM *ULMap) PositionAddressInfoUpdate(ul models.Positionaddressinfo) error {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	ULM.PositionAddressInfoMap[ul.Usercode] = ul

	return nil
}

func (ULM *ULMap) PositionAddressInfoDelete(ul models.Positionaddressinfo) error {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	delete(ULM.PositionAddressInfoMap, ul.Usercode)

	return nil
}
