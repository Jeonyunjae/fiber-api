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
	PositionAddressInfosInit() error
	PositionAddressInfosInsert(models.PositionAddressInfo) error
	PositionAddressInfosRead(models.PositionAddressInfo) ([]models.PositionAddressInfo, error)
	PositionAddressInfosAllRead() ([]models.PositionAddressInfo, error)
	PositionAddressInfosUpdate(models.PositionAddressInfo) (bool, error)
	PositionAddressInfosDelete(models.PositionAddressInfo) (bool, error)
}

var IORM ICrud

type ULStuct struct {
	PositionAddressInfoMap map[int]models.PositionAddressInfo
}

// PositionAddressInfo Data 정의
func ServiceInit() error {

	// 1.orm
	err := myorm.PositionAddressInfo.PositionAddressInfosInit()
	if err != nil {
		return err
	}

	// 2.query
	err = myquery.PositionAddressInfo.PositionAddressInfosInit()
	if err != nil {
		return err
	}

	tempMapeData, err := myquery.PositionAddressInfo.PositionAddressInfoAllRead()
	if tempMapeData == nil || err != nil {
		return err
	}

	// 3.slice
	err = myslice.PositionAddressInfo.PositionAddressInfosInit(
		share.PositionAddressInfoMapToSlice(tempMapeData))
	if err != nil {
		return err
	}

	// 4.map
	err = mymap.PositionAddressInfo.PositionAddressInfosInit(tempMapeData)
	if err != nil {
		return err
	}

	// 5.kdtree
	err = mykdtree.PositionAddressInfo.PositionAddressInfosInit(tempMapeData)
	if err != nil {
		return err
	}

	return nil
}

// PositionAddressInfo 데이터 추가
func ServiceInsert(PositionAddressInfo models.PositionAddressInfo) error {

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

// PositionAddressInfo Data 특정 하나 가져오기
func ServiceRead(PositionAddressInfo models.PositionAddressInfo) error {
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
		//myorm.PositionAddressInfo.PositionAddressInfoRead(PositionAddressInfo)
	}()

	// 6.query
	go func() {
		defer wg.Done()
		myquery.PositionAddressInfo.PositionAddressInfoRead(PositionAddressInfo)
	}()

	wg.Wait()

	return nil
}

// PositionAddressInfo Data 특정 조건의 값들 가져오기
func ServiceReads(PositionAddressInfo models.PositionAddressInfo) error {
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
