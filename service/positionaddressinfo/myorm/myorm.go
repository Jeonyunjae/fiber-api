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

func (ULO *ULOrm) PositionAddressInfoInit() error {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	if mydbgorm.Database.Db == nil {
		return log.MyError("Error_PositionAddressInfoInit")
	}
	PositionAddressInfo.PositionAddressInfoOrm = mydbgorm.Database.Db

	return nil
}

func (ULO *ULOrm) PositionAddressInfoInsert(ul models.Positionaddressinfo) error {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	err := ULO.PositionAddressInfoOrm.Create(ul)
	if err == nil {
		return err.Error
	}
	return nil
}

func (ULO *ULOrm) PositionAddressInfoRead(ul models.Positionaddressinfo) ([]models.Positionaddressinfo, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	var rows []models.Positionaddressinfo
	var row models.Positionaddressinfo
	ULO.PositionAddressInfoOrm.First(&row, "usercode = ?", ul.Usercode) // primary key기준으로 product 찾기
	rows = append(rows, row)

	return rows, nil
}

func (ULO *ULOrm) PositionAddressInfoAllRead() ([]models.Positionaddressinfo, error) {
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

func (ULO *ULOrm) PositionAddressInfoUpdate(ul models.Positionaddressinfo) (bool, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	return false, nil
}

func (ULO *ULOrm) PositionAddressInfoDelete(ul models.Positionaddressinfo) (bool, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	return false, nil
}
