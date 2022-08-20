package myorm

import (
	"github.com/jeonyunjae/fiber-api/database/mydbgorm"
	"github.com/jeonyunjae/fiber-api/models"
	"github.com/jeonyunjae/fiber-api/util/log"
	"gorm.io/gorm"
)

var PositionAddressInfo ULOrm

type ULOrm struct {
	PositionAddressInfoOrm *gorm.DB
}

func (ULO *ULOrm) PositionAddressInfosInit() error {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	if mydbgorm.Database.Db == nil {
		return log.MyError("Error_PositionAddressInfosInit")
	}
	PositionAddressInfo.PositionAddressInfoOrm = mydbgorm.Database.Db

	return nil
}

func (ULO *ULOrm) PositionAddressInfoInsert(ul models.PositionAddressInfo) error {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	err := ULO.PositionAddressInfoOrm.Create(ul)
	if err == nil {
		return err.Error
	}
	return nil
}

func (ULO *ULOrm) PositionAddressInfoRead(ul models.PositionAddressInfo) ([]models.PositionAddressInfo, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	// works with Take
	result := map[models.PositionAddressInfo]interface{}{}
	ULO.PositionAddressInfoOrm.Table("PositionAddressInfo").Take(&result)

	return nil, nil
}

func (ULO *ULOrm) PositionAddressInfoAllRead() ([]models.PositionAddressInfo, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	// works with Take
	//result := map[models.PositionAddressInfo]interface{}{}
	//ULO.PositionAddressInfoOrm.Table("PositionAddressInfo").Take(&result)

	data := ULO.PositionAddressInfoOrm.Select("UserCode")
	if data == nil {
		return nil, nil
	}

	return nil, nil
}

func (ULO *ULOrm) PositionAddressInfoUpdate(ul models.PositionAddressInfo) (bool, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	return false, nil
}

func (ULO *ULOrm) PositionAddressInfoDelete(ul models.PositionAddressInfo) (bool, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	return false, nil
}
