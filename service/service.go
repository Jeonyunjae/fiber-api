package service

import (
	"github.com/jeonyunjae/fiber-api/models"
	"github.com/jeonyunjae/fiber-api/service/positionaddressinfo/mykdtree"
	"github.com/jeonyunjae/fiber-api/service/positionaddressinfo/mymap"
	"github.com/jeonyunjae/fiber-api/service/positionaddressinfo/myorm"
	"github.com/jeonyunjae/fiber-api/service/positionaddressinfo/myquery"
	"github.com/jeonyunjae/fiber-api/service/positionaddressinfo/myslice"
	"github.com/jeonyunjae/fiber-api/util/log"
)

type ULStuct struct {
	PositionAddressInfoMap map[int]models.PositionAddressInfo
}

// PositionAddressInfo Data 정의
func ServiceInit() error {

	// 1.orm
	ormData := myorm.PositionAddressInfo.PositionAddressInfosInit()
	if ormData.PositionAddressInfoOrm == nil {
		return log.MyError("Error_ServiceInit_Orm")
	}

	tempSliceData, err := myorm.PositionAddressInfo.PositionAddressInfoAllRead()
	if tempSliceData == nil || err != nil {
		return log.MyError("Error_ServiceInit_ormData.PositionAddressInfoAllRead")
	}
	// 2.slice
	sliceData := myslice.PositionAddressInfo.PositionAddressInfosInit()
	if sliceData.PositionAddressInfoSlice == nil {
		return log.MyError("Error_ServiceInit_Slice")
	}

	// 3.map
	mapData := mymap.PositionAddressInfo.PositionAddressInfosInit()
	if mapData.PositionAddressInfoMap == nil {
		return log.MyError("Error_ServiceInit_Map")
	}

	// 4.kdtree
	kdtreeData := mykdtree.PositionAddressInfo.PositionAddressInfosInit()
	if kdtreeData.PositionAddressInfoKdTree.Points() == nil {
		return log.MyError("Error_ServiceInit_KdTree")
	}

	// 5.query
	queryData := myquery.PositionAddressInfo.PositionAddressInfosInit()
	if queryData.PositionAddressInfoQuery == nil {
		return log.MyError("Error_ServiceInit_Query")
	}
	//test
	//util.PostAddressDefineCsvToDb()
	return nil
}

//PositionAddressInfo 데이터 추가
func ServiceInsert(PositionAddressInfo models.PositionAddressInfo) error {
	// 1.slice
	resSlice := myslice.PositionAddressInfo.PositionAddressInfoInsert(PositionAddressInfo)
	if resSlice.PositionAddressInfoSlice == nil {
		return log.MyError("CreateListError")
	}
	// 2.decimaltree

	//3.map
	resMap := mymap.PositionAddressInfo.PositionAddressInfoInsert(PositionAddressInfo)
	if resMap.PositionAddressInfoMap == nil {
		return log.MyError("CreateListError")
	}
	// 4.kdtree
	resKDTree := mykdtree.PositionAddressInfo.PositionAddressInfoInsert(PositionAddressInfo)
	resKDTree.PositionAddressInfoKdTree.Balance()

	// 5.orm
	resOrm := myorm.PositionAddressInfo.PositionAddressInfoInsert(PositionAddressInfo)
	if resOrm == nil {
		return log.MyError("CreateListError" + resOrm.Error())
	}
	// 6.query
	resQuery := myquery.PositionAddressInfo.PositionAddressInfoInsert(PositionAddressInfo)
	if resQuery == nil {
		return log.MyError("CreateListError" + resQuery.Error())
	}

	return nil
}

// PositionAddressInfo Data 전체 가져오기
func ServiceRead(PositionAddressInfo models.PositionAddressInfo) error {

	// 1.slice
	myslice.PositionAddressInfo.PositionAddressInfoRead(PositionAddressInfo)

	// 2.decimaltree

	//3.map
	mymap.PositionAddressInfo.PositionAddressInfoRead(PositionAddressInfo)

	// 4.kdtree
	mykdtree.PositionAddressInfo.PositionAddressInfoRead(PositionAddressInfo)

	// 5.orm
	myorm.PositionAddressInfo.PositionAddressInfoRead(PositionAddressInfo)

	// 6.query
	myquery.PositionAddressInfo.PositionAddressInfoRead(PositionAddressInfo)

	return nil
}

func ServiceUpdate(PositionAddressInfo models.PositionAddressInfo) error {

	// 1.slice
	myslice.PositionAddressInfo.PositionAddressInfoUpdate(PositionAddressInfo)

	// 2.decimaltree

	//3.map
	mymap.PositionAddressInfo.PositionAddressInfoUpdate(PositionAddressInfo)

	// 4.kdtree
	mykdtree.PositionAddressInfo.PositionAddressInfoUpdate(PositionAddressInfo)

	// 5.orm
	myorm.PositionAddressInfo.PositionAddressInfoUpdate(PositionAddressInfo)

	// 6.query
	myquery.PositionAddressInfo.PositionAddressInfoUpdate(PositionAddressInfo)

	return nil
}

func ServiceDelete(PositionAddressInfo models.PositionAddressInfo) error {

	// 1.slice
	myslice.PositionAddressInfo.PositionAddressInfoDelete(PositionAddressInfo)

	// 2.decimaltree

	//3.map
	mymap.PositionAddressInfo.PositionAddressInfoDelete(PositionAddressInfo)

	// 4.kdtree
	mykdtree.PositionAddressInfo.PositionAddressInfoDelete(PositionAddressInfo)

	// 5.orm
	myorm.PositionAddressInfo.PositionAddressInfoDelete(PositionAddressInfo)

	// 6.query
	myquery.PositionAddressInfo.PositionAddressInfoDelete(PositionAddressInfo)

	return nil
}
