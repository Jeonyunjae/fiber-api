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

func (ULKDT *ULKDTree) PositionAddressInfoInit(rows map[string]models.Positionaddressinfo) error {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	if rows == nil {
		return log.MyError("Error_PositionAddressInfoInit")
	}

	PositionAddressInfo = PositionAddressInfo.PositionAddressInfoMapToKDTree(rows)
	return nil
}

func (ULKDT *ULKDTree) PositionAddressInfoArrayToKDTree(data [][]string) ULKDTree {
	for _, row := range data {
		var PositionAddressInfo models.Positionaddressinfo
		PositionAddressInfo.Usercode = row[0]
		PositionAddressInfo.Loclatitude, _ = strconv.ParseFloat(row[2], 64)
		PositionAddressInfo.Loclongtitude, _ = strconv.ParseFloat(row[3], 64)
		ULKDT.PositionAddressInfoKdTree.Insert(
			points.NewPoint(
				[]float64{PositionAddressInfo.Loclongtitude, PositionAddressInfo.Loclatitude}, PositionAddressInfo))
	}

	return *ULKDT
}

func (ULKDT *ULKDTree) PositionAddressInfoMapToKDTree(rows map[string]models.Positionaddressinfo) ULKDTree {
	for _, row := range rows {

		point := ULKDT.PositionAddressInfoKdTree.Find(points.NewPoint(
			[]float64{row.Loclongtitude, row.Loclatitude}, row))

		if point == nil {
			ULKDT.PositionAddressInfoKdTree.Insert(
				points.NewPoint(
					[]float64{row.Loclongtitude, row.Loclatitude}, row))
		}
	}
	return *ULKDT
}

func (ULKDT *ULKDTree) PositionAddressInfoInsert(ul models.Positionaddressinfo) error {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	ULKDT.PositionAddressInfoKdTree.Insert(
		points.NewPoint(
			[]float64{ul.Loclongtitude, ul.Loclatitude}, ul))

	data, err := ULKDT.PositionAddressInfoRead(ul)
	if len(data) < 1 || err != nil {
		return err
	}
	return nil
}

func (ULKDT *ULKDTree) PositionAddressInfoRead(ul models.Positionaddressinfo) ([]models.Positionaddressinfo, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	var result []models.Positionaddressinfo
	nodes := ULKDT.PositionAddressInfoKdTree.Find(&points.Point2D{Y: ul.Loclatitude, X: ul.Loclongtitude})

	if nodes == nil {
		return nil, log.MyError("Error_PositionAddressInfoRead")
	}
	value := nodes.(*points.Point).Data.(models.Positionaddressinfo)
	result = append(result, models.Positionaddressinfo{Usercode: value.Usercode, Loclongtitude: value.Loclongtitude, Loclatitude: value.Loclatitude})

	return result, nil
}

func (ULKDT *ULKDTree) PositionAddressInfoUpdate(PositionAddressInfo models.Positionaddressinfo) ULKDTree {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	return *ULKDT
}

func (ULKDT *ULKDTree) PositionAddressInfoDelete(PositionAddressInfo models.Positionaddressinfo) ULKDTree {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	return *ULKDT
}
