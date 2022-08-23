package service

import (
	"sync"

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

// PositionAddressInfo Data 특정 조건의 값들 가져오기
func ServiceReads(PositionAddressInfo models.Positionaddressinfo) error {
	var wg = sync.WaitGroup{}
	wg.Add(6)

	// 1.slice
	go func() {
		defer wg.Done()
		myslice.PositionAddressInfo.PositionAddressInfoRead(PositionAddressInfo)
	}()
	// 2.decimaltree

	//3.map
	go func() {
		defer wg.Done()
		mymap.PositionAddressInfo.PositionAddressInfoRead(PositionAddressInfo)

	}()
	// 4.kdtree
	go func() {
		defer wg.Done()
		mykdtree.PositionAddressInfo.PositionAddressInfoRead(PositionAddressInfo)
	}()

	// 5.orm
	go func() {
		defer wg.Done()
		myorm.PositionAddressInfo.PositionAddressInfoRead(PositionAddressInfo)
	}()

	// 6.query
	go func() {
		defer wg.Done()
		myquery.PositionAddressInfo.PositionAddressInfoRead(PositionAddressInfo)
	}()

	wg.Wait()

	return nil
}

func ServiceUpdate(PositionAddressInfo models.Positionaddressinfo) error {
	// 1.slice
	result, err := myslice.PositionAddressInfo.PositionAddressInfoUpdate(PositionAddressInfo)
	if err != nil || result == false {
		return err
	}

	// 2.map
	result, err = mymap.PositionAddressInfo.PositionAddressInfoUpdate(PositionAddressInfo)
	if err != nil || result == false {
		return err
	}

	// 3.kdtree
	result, err = mykdtree.PositionAddressInfo.PositionAddressInfoUpdate(PositionAddressInfo)
	if err != nil || result == false {
		return err
	}

	// 4.orm
	result, err = myorm.PositionAddressInfo.PositionAddressInfoUpdate(PositionAddressInfo)
	if err != nil || result == false {
		return err
	}

	// 5.query
	result, err = myquery.PositionAddressInfo.PositionAddressInfoUpdate(PositionAddressInfo)
	if err != nil || result == false {
		return err
	}

	return nil
}

func ServiceDelete(PositionAddressInfo models.Positionaddressinfo) error {
	// 1.slice
	result, err := myslice.PositionAddressInfo.PositionAddressInfoDelete(PositionAddressInfo)
	if err != nil || result == false {
		return err
	}

	// 2.map
	result, err = mymap.PositionAddressInfo.PositionAddressInfoDelete(PositionAddressInfo)
	if err != nil || result == false {
		return err
	}

	// 3.kdtree
	result, err = mykdtree.PositionAddressInfo.PositionAddressInfoDelete(PositionAddressInfo)
	if err != nil || result == false {
		return err
	}

	// 4.orm
	result, err = myorm.PositionAddressInfo.PositionAddressInfoDelete(PositionAddressInfo)
	if err != nil || result == false {
		return err
	}

	// 5.query
	result, err = myquery.PositionAddressInfo.PositionAddressInfoDelete(PositionAddressInfo)
	if err != nil || result == false {
		return err
	}

	return nil
}
