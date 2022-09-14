package myorm

import (
	"github.com/jeonyunjae/fiber-api/datatype/gorm"
	"github.com/jeonyunjae/fiber-api/models"
	"github.com/jeonyunjae/fiber-api/util/log"
)

var PositionAddressInfo ULOrm

type ULOrm struct {
	PositionAddressInfoOrm gorm.DBInstance
}

func (ULO *ULOrm) PositionAddressInfoInit() error {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	if gorm.Database.Db == nil {
		return log.MyError("Error_PositionAddressInfoInit")
	}
	PositionAddressInfo.PositionAddressInfoOrm = gorm.Database

	return nil
}

func (ULO *ULOrm) PositionAddressInfoInsert(ul models.Positionaddressinfo) error {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	err := ULO.PositionAddressInfoOrm.Db.Create(ul)
	if err.Error != nil {
		return err.Error
	}

	data, err_Read := ULO.PositionAddressInfoRead(ul)
	if len(data) < 1 || err_Read != nil {
		return err_Read
	}
	return nil
}

func (ULO *ULOrm) PositionAddressInfoRead(ul models.Positionaddressinfo) (map[string]models.Positionaddressinfo, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	rows := make(map[string]models.Positionaddressinfo)
	var row models.Positionaddressinfo

	ULO.PositionAddressInfoOrm.Db.First(&row, "usercode = ?", ul.Usercode) // primary key기준으로 product 찾기
	rows[row.Usercode] = row

	return rows, nil
}

func (ULO *ULOrm) PositionAddressInfoAllRead() ([]models.Positionaddressinfo, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	data := ULO.PositionAddressInfoOrm.Db.Select("UserCode")
	if data.Error != nil {
		return nil, log.MyError(data.Error.Error())
	}

	return nil, nil
}

func (ULO *ULOrm) PositionAddressInfoUpdate(ul models.Positionaddressinfo) error {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	data := ULO.PositionAddressInfoOrm.Db.Model(&models.Positionaddressinfo{}).Where("usercode = ?", ul.Usercode).Updates(ul)
	if data.Error != nil {
		return log.MyError(data.Error.Error())
	}
	return nil
}

func (ULO *ULOrm) PositionAddressInfoDelete(ul models.Positionaddressinfo) error {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	data := ULO.PositionAddressInfoOrm.Db.Where("usercode = ?", ul.Usercode).Delete(&models.Positionaddressinfo{})
	if data.Error != nil {
		return log.MyError(data.Error.Error())
	}

	return nil
}
