package mykdtree

import (
	"strconv"

	"github.com/jeonyunjae/fiber-api/kdtree"
	"github.com/jeonyunjae/fiber-api/kdtree/points"
	"github.com/jeonyunjae/fiber-api/models"
	"github.com/jeonyunjae/fiber-api/util/log"
)

var PositionAddressInfo ULKDTree

type ULKDTree struct {
	PositionAddressInfoKdTree kdtree.KDTree
}

func (ULKDT *ULKDTree) PositionAddressInfosInit(rows map[string]models.PositionAddressInfo) error {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	if rows == nil {
		return log.MyError("Error_PositionAddressInfosInit")
	}

	PositionAddressInfo = PositionAddressInfo.PositionAddressInfoMapToKDTree(rows)
	return nil
}

func (ULKDT *ULKDTree) PositionAddressInfoArrayToKDTree(data [][]string) ULKDTree {
	for _, row := range data {
		var PositionAddressInfo models.PositionAddressInfo
		PositionAddressInfo.UserCode = row[0]
		PositionAddressInfo.LocLatitude, _ = strconv.ParseFloat(row[2], 64)
		PositionAddressInfo.LocLongtitude, _ = strconv.ParseFloat(row[3], 64)
		ULKDT.PositionAddressInfoKdTree.Insert(
			points.NewPoint(
				[]float64{PositionAddressInfo.LocLongtitude, PositionAddressInfo.LocLatitude, 0}, PositionAddressInfo))
	}

	return *ULKDT
}

func (ULKDT *ULKDTree) PositionAddressInfoMapToKDTree(rows map[string]models.PositionAddressInfo) ULKDTree {
	for _, row := range rows {
		ULKDT.PositionAddressInfoKdTree.Insert(
			points.NewPoint(
				[]float64{row.LocLongtitude, row.LocLatitude, 0}, row))

	}

	return *ULKDT
}

func (ULKDT *ULKDTree) PositionAddressInfoInsert(ul models.PositionAddressInfo) error {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	ULKDT.PositionAddressInfoKdTree.Insert(
		points.NewPoint(
			[]float64{ul.LocLongtitude, ul.LocLatitude, 0}, ul))

	data, err := ULKDT.PositionAddressInfoRead(ul)
	if data.UserCode == "" || err != nil {
		return err
	}
	return nil
}

func (ULKDT *ULKDTree) PositionAddressInfoRead(PositionAddressInfo models.PositionAddressInfo) (models.PositionAddressInfo, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	var data = models.PositionAddressInfo{UserCode: "11111", LocLatitude: 1, LocLongtitude: 1}
	ULKDT.PositionAddressInfoKdTree.Find(&points.Point2D{X: 36.458658, Y: 128.891228})

	return data, nil
}

func (ULKDT *ULKDTree) PositionAddressInfoUpdate(PositionAddressInfo models.PositionAddressInfo) ULKDTree {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	return *ULKDT
}

func (ULKDT *ULKDTree) PositionAddressInfoDelete(PositionAddressInfo models.PositionAddressInfo) ULKDTree {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	return *ULKDT
}
