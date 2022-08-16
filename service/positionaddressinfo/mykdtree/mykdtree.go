package mykdtree

import (
	"strconv"

	"github.com/jeonyunjae/fiber-api/models"
	"github.com/jeonyunjae/fiber-api/service/util"
	"github.com/jeonyunjae/fiber-api/util/log"
	"github.com/kyroy/kdtree"
	"github.com/kyroy/kdtree/points"
)

var PositionAddressInfo ULKDTree

type ULKDTree struct {
	PositionAddressInfoKdTree kdtree.KDTree
}

func (ULKDT *ULKDTree) PositionAddressInfosInit() ULKDTree {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	data := util.PositionAddressInfoCsvToSlice()

	PositionAddressInfo = PositionAddressInfo.PositionAddressInfoArrayToKDTree(data)
	return *ULKDT
}

func (ULKDT *ULKDTree) PositionAddressInfoArrayToKDTree(data [][]string) ULKDTree {
	for _, row := range data {
		var PositionAddressInfo models.PositionAddressInfo
		PositionAddressInfo.UserCode, _ = strconv.ParseUint(row[0], 10, 32)
		PositionAddressInfo.LocLatitude, _ = strconv.ParseFloat(row[2], 64)
		PositionAddressInfo.LocLongtitue, _ = strconv.ParseFloat(row[3], 64)
		ULKDT.PositionAddressInfoKdTree.Insert(
			points.NewPoint(
				[]float64{PositionAddressInfo.LocLongtitue, PositionAddressInfo.LocLatitude, 0}, PositionAddressInfo))
	}

	return *ULKDT
}

func (ULKDT *ULKDTree) PositionAddressInfoInsert(PositionAddressInfo models.PositionAddressInfo) ULKDTree {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	ULKDT.PositionAddressInfoKdTree.Insert(
		points.NewPoint(
			[]float64{PositionAddressInfo.LocLongtitue, PositionAddressInfo.LocLatitude, 0}, PositionAddressInfo))

	return *ULKDT
}

func (ULKDT *ULKDTree) PositionAddressInfoRead(PositionAddressInfo models.PositionAddressInfo) ULKDTree {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	return *ULKDT
}

func (ULKDT *ULKDTree) PositionAddressInfoUpdate(PositionAddressInfo models.PositionAddressInfo) ULKDTree {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	return *ULKDT
}

func (ULKDT *ULKDTree) PositionAddressInfoDelete(PositionAddressInfo models.PositionAddressInfo) ULKDTree {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	return *ULKDT
}
