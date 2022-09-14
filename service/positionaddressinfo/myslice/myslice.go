package myslice

import (
	"sort"

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

func (ULL *ULSlice) PositionAddressInfoRead(ul models.Positionaddressinfo) (map[string]models.Positionaddressinfo, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	rows := make(map[string]models.Positionaddressinfo)
	var row models.Positionaddressinfo

	for _, row = range ULL.PositionAddressInfoSlice {
		if row.Usercode == ul.Usercode {
			row.Loclatitude = ul.Loclatitude
			row.Loclongtitude = ul.Loclongtitude
			rows[row.Usercode] = row
			return rows, nil
		}
	}
	return rows, log.MyError("Error_PositionAddressInfoRead")
}

func (ULL *ULSlice) PositionAddressInfoReads(ul models.PositionaddressDistanceInfo) ([]models.PositionaddressDistanceInfo, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	var sortData []models.PositionaddressDistanceInfo

	for _, row := range ULL.PositionAddressInfoSlice {
		var data models.PositionaddressDistanceInfo
		distance := util.GetDistance(ul.Loclongtitude, row.Loclongtitude, ul.Loclatitude, row.Loclatitude)

		data.Usercode = row.Usercode
		data.Loclongtitude = row.Loclongtitude
		data.Loclatitude = row.Loclatitude
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

func (ULL *ULSlice) PositionAddressInfoReadsRange(ul models.PositionaddressDistanceInfo) ([]models.PositionaddressDistanceInfo, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	var sortData []models.PositionaddressDistanceInfo

	for _, row := range ULL.PositionAddressInfoSlice {
		distance := util.GetDistance(ul.Loclongtitude, row.Loclongtitude, ul.Loclatitude, row.Loclatitude)
		if distance <= ul.Distance {
			sortData = append(sortData, models.PositionaddressDistanceInfo{Usercode: row.Usercode, Loclatitude: row.Loclatitude, Loclongtitude: row.Loclongtitude, Distance: distance})
		}
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

func (ULL *ULSlice) PositionAddressInfoUpdate(ul models.Positionaddressinfo) error {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	var row models.Positionaddressinfo
	for _, row = range ULL.PositionAddressInfoSlice {
		if row.Usercode == ul.Usercode {
			row.Loclatitude = ul.Loclatitude
			row.Loclongtitude = ul.Loclongtitude
			return nil
		}
	}
	return log.MyError("NotFound")
}

func (ULL *ULSlice) PositionAddressInfoDelete(ul models.Positionaddressinfo) error {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	var index int
	var row models.Positionaddressinfo
	for index, row = range ULL.PositionAddressInfoSlice {
		if row.Usercode == ul.Usercode {
			break
		}
	}

	ULL.PositionAddressInfoSlice = ULL.sliceDeelte(index)
	return nil
}

func (ULL *ULSlice) sliceDeelte(index int) []models.Positionaddressinfo {
	return append(ULL.PositionAddressInfoSlice[:index], ULL.PositionAddressInfoSlice[index+1:]...)
}
