package service

import (
	"github.com/jeonyunjae/fiber-api/models"
	"github.com/jeonyunjae/fiber-api/service/positionaddressinfo/mykdtree"
	"github.com/jeonyunjae/fiber-api/service/positionaddressinfo/mymap"
	"github.com/jeonyunjae/fiber-api/service/positionaddressinfo/myorm"
	"github.com/jeonyunjae/fiber-api/service/positionaddressinfo/myquery"
	"github.com/jeonyunjae/fiber-api/service/positionaddressinfo/myslice"
	"github.com/jeonyunjae/fiber-api/service/positionaddressinfo/share"
)

type ICrud interface {
	PositionAddressInfoInit() error
	PositionAddressInfosInsert(models.Positionaddressinfo) error
	PositionAddressInfosRead(models.Positionaddressinfo) ([]models.Positionaddressinfo, error)
	PositionAddressInfosAllRead() ([]models.Positionaddressinfo, error)
	PositionAddressInfosUpdate(models.Positionaddressinfo) (bool, error)
	PositionAddressInfosDelete(models.Positionaddressinfo) (bool, error)
}

var IORM ICrud

type ULStuct struct {
	PositionAddressInfoMap map[int]models.Positionaddressinfo
}

// PositionAddressInfo Data 정의
func ServiceInit() error {

	// 1.orm
	err := myorm.PositionAddressInfo.PositionAddressInfoInit()
	if err != nil {
		return err
	}

	// 2.query
	err = myquery.PositionAddressInfo.PositionAddressInfoInit()
	if err != nil {
		return err
	}

	tempMapeData, err := myquery.PositionAddressInfo.PositionAddressInfoAllRead()
	if tempMapeData == nil || err != nil {
		return err
	}

	// 3.slice
	err = myslice.PositionAddressInfo.PositionAddressInfoInit(
		share.PositionAddressInfoMapToSlice(tempMapeData))
	if err != nil {
		return err
	}

	// 4.map
	err = mymap.PositionAddressInfo.PositionAddressInfoInit(tempMapeData)
	if err != nil {
		return err
	}

	// 5.kdtree
	err = mykdtree.PositionAddressInfo.PositionAddressInfoInit(tempMapeData)
	if err != nil {
		return err
	}

	return nil
}

// PositionAddressInfo 데이터 추가
func ServiceInsert(PositionAddressInfo models.Positionaddressinfo) error {

	// 1.slice
	err := myslice.PositionAddressInfo.PositionAddressInfoInsert(PositionAddressInfo)
	if err != nil {
		return err
	}

	//2.map
	err = mymap.PositionAddressInfo.PositionAddressInfoInsert(PositionAddressInfo)
	if err != nil {
		return err
	}

	// 3.kdtree
	err = mykdtree.PositionAddressInfo.PositionAddressInfoInsert(PositionAddressInfo)
	if err != nil {
		return err
	}

	// 4.orm
	err = myorm.PositionAddressInfo.PositionAddressInfoInsert(PositionAddressInfo)
	if err != nil {
		return err
	}

	// 5.query
	err = myquery.PositionAddressInfo.PositionAddressInfoInsert(PositionAddressInfo)
	if err != nil {
		return err
	}

	return nil
}

//usercode로 정보 가져오기
func ServiceRead(PositionAddressInfo models.Positionaddressinfo) error {

	// 1.slice
	result, err := myslice.PositionAddressInfo.PositionAddressInfoRead(PositionAddressInfo)
	if err != nil || len(result) < 1 {
		return err
	}

	// 2.map
	result, err = mymap.PositionAddressInfo.PositionAddressInfoRead(PositionAddressInfo)
	if err != nil || len(result) < 1 {
		return err
	}

	// 3.kdtree
	result, err = mykdtree.PositionAddressInfo.PositionAddressInfoRead(PositionAddressInfo)
	if err != nil || len(result) < 1 {
		return err
	}

	// 4.orm
	result, err = myorm.PositionAddressInfo.PositionAddressInfoRead(PositionAddressInfo)
	if err != nil || len(result) < 1 {
		return err
	}

	// 5.query
	result, err = myquery.PositionAddressInfo.PositionAddressInfoRead(PositionAddressInfo)
	if err != nil || len(result) < 1 {
		return err
	}

	return nil
}

// PositionAddressInfo 가까이 있는 정보 가져오기 특정 갯수 가져오기
func ServiceReadsLimit(PositionAddressDistanceInfo models.PositionaddressDistanceInfo) error {
	// 1.slice
	result, err := myslice.PositionAddressInfo.PositionAddressInfoReads(PositionAddressDistanceInfo)
	if err != nil || len(result) < 1 {
		return err
	}

	// 2.map
	result, err = mymap.PositionAddressInfo.PositionAddressInfoReads(PositionAddressDistanceInfo)
	if err != nil || len(result) < 1 {
		return err
	}

	// 3.kdtree
	result, err = mykdtree.PositionAddressInfo.PositionAddressInfoReads(PositionAddressDistanceInfo)
	if err != nil || len(result) < 1 {
		return err
	}

	// // 4.orm
	// result, err = myorm.PositionAddressInfo.PositionAddressInfoRead(PositionAddressDistanceInfo)
	// if err != nil || len(result) < 1 {
	// 	return err
	// }

	// 5.query
	result, err = myquery.PositionAddressInfo.PositionAddressInfoReads(PositionAddressDistanceInfo)
	if err != nil || len(result) < 1 {
		return err
	}

	return nil
}

func ServiceUpdate(PositionAddressInfo models.Positionaddressinfo) error {
	// 1.slice
	err := myslice.PositionAddressInfo.PositionAddressInfoUpdate(PositionAddressInfo)
	if err != nil {
		return err
	}

	// 2.map
	err = mymap.PositionAddressInfo.PositionAddressInfoUpdate(PositionAddressInfo)
	if err != nil {
		return err
	}

	// 3.kdtree
	err = mykdtree.PositionAddressInfo.PositionAddressInfoUpdate(PositionAddressInfo)
	if err != nil {
		return err
	}

	// 4.orm
	err = myorm.PositionAddressInfo.PositionAddressInfoUpdate(PositionAddressInfo)
	if err != nil {
		return err
	}

	// 5.query
	err = myquery.PositionAddressInfo.PositionAddressInfoUpdate(PositionAddressInfo)
	if err != nil {
		return err
	}

	return nil
}

func ServiceDelete(PositionAddressInfo models.Positionaddressinfo) error {
	// 1.slice
	err := myslice.PositionAddressInfo.PositionAddressInfoDelete(PositionAddressInfo)
	if err != nil {
		return err
	}

	// 2.map
	err = mymap.PositionAddressInfo.PositionAddressInfoDelete(PositionAddressInfo)
	if err != nil {
		return err
	}

	// 3.kdtree
	err = mykdtree.PositionAddressInfo.PositionAddressInfoDelete(PositionAddressInfo)
	if err != nil {
		return err
	}

	// 4.orm
	err = myorm.PositionAddressInfo.PositionAddressInfoDelete(PositionAddressInfo)
	if err != nil {
		return err
	}

	// 5.query
	err = myquery.PositionAddressInfo.PositionAddressInfoDelete(PositionAddressInfo)
	if err != nil {
		return err
	}

	return nil
}
