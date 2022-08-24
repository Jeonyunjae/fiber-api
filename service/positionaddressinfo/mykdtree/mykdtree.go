package mykdtree

import (
	"fmt"

	"github.com/jeonyunjae/fiber-api/datatype/kdtree"
	"github.com/jeonyunjae/fiber-api/datatype/kdtree/points"
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

func (ULKDT *ULKDTree) PositionAddressInfoMapToKDTree(rows map[string]models.Positionaddressinfo) ULKDTree {
	var num int
	for _, row := range rows {
		var inRes []models.Positionaddressinfo
		point := ULKDT.PositionAddressInfoKdTree.Find(points.NewPoint(
			[]float64{row.Loclongtitude, row.Loclatitude}, row))

		if point != nil {
			inRes = point.(*points.Point).Data.([]models.Positionaddressinfo)
			inRes = append(inRes, row)
		} else {
			inRes = append(inRes, row)
		}

		ULKDT.PositionAddressInfoKdTree.Insert(
			points.NewPoint(
				[]float64{row.Loclongtitude, row.Loclatitude}, inRes))

		log.MyLog(fmt.Sprintf("%d==%s", num, row.Usercode))
		num = num +1
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

func (ULKDT *ULKDTree) PositionAddressInfoRead(ul models.Positionaddressinfo) (map[string]models.Positionaddressinfo, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	rows := make(map[string]models.Positionaddressinfo)
	nodes := ULKDT.PositionAddressInfoKdTree.Find(&points.Point2D{Y: ul.Loclatitude, X: ul.Loclongtitude})

	if nodes == nil {
		return nil, log.MyError("Error_PositionAddressInfoRead")
	}

	value := nodes.(*points.Point).Data.(models.Positionaddressinfo)

	rows[value.Usercode] = models.Positionaddressinfo{Usercode: value.Usercode, Loclongtitude: value.Loclongtitude, Loclatitude: value.Loclatitude}

	return rows, nil
}

func (ULKDT *ULKDTree) PositionAddressInfoReads(ul models.PositionaddressDistanceInfo) ([]models.PositionaddressDistanceInfo, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	var rows []models.PositionaddressDistanceInfo
	nodes := ULKDT.PositionAddressInfoKdTree.KNN(&points.Point{Coordinates: []float64{ul.Loclatitude, ul.Loclongtitude}}, 100)

	if nodes == nil {
		return nil, log.MyError("Error_PositionAddressInfoRead")
	}
	var info models.PositionaddressDistanceInfo
	var infoTemp models.Positionaddressinfo

	nodeLen := len(nodes)
	for i := 0; i < nodeLen; i++ {
		nodeInLen := len(nodes[i].(*points.Point).Data.([]models.Positionaddressinfo))

		for j := 0; j < nodeInLen; j++ {
			infoTemp = nodes[i].(*points.Point).Data.([]models.Positionaddressinfo)[j]
			info.Usercode = infoTemp.Usercode
			info.Loclatitude = infoTemp.Loclatitude
			info.Loclongtitude = infoTemp.Loclongtitude

			rows = append(rows, info)
		}
	}
	//rows[value.Usercode] = models.Positionaddressinfo{Usercode: value.Usercode, Loclongtitude: value.Loclongtitude, Loclatitude: value.Loclatitude}

	return rows, nil
}

func (ULKDT *ULKDTree) PositionAddressInfoUpdate(ul models.Positionaddressinfo) (bool, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	_, err := ULKDT.PositionAddressInfoDelete(ul)

	if err != nil {
		return false, log.MyError("Error_PositionAddressInfoUpdate")
	}

	err = ULKDT.PositionAddressInfoInsert(ul)
	if err != nil {
		return false, log.MyError("Error_PositionAddressInfoUpdate")
	}

	return true, nil
}

func (ULKDT *ULKDTree) PositionAddressInfoDelete(ul models.Positionaddressinfo) (bool, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	ULKDT.PositionAddressInfoKdTree.Remove(points.NewPoint(
		[]float64{ul.Loclongtitude, ul.Loclatitude}, ul))

	return true, nil
}
