package myquery

import (
	"database/sql"

	"github.com/jeonyunjae/fiber-api/database/mydbquery"
	"github.com/jeonyunjae/fiber-api/models"
	"github.com/jeonyunjae/fiber-api/util/log"
)

var PositionAddressInfo ULQuery

type ULQuery struct {
	PositionAddressInfoQuery *sql.DB
}

func (ULQ *ULQuery) PositionAddressInfosInit() ULQuery {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	PositionAddressInfo.PositionAddressInfoQuery = mydbquery.Database.Db

	return *ULQ
}

func (ULQ *ULQuery) PositionAddressInfoInsert(PositionAddressInfo models.PositionAddressInfo) error {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	resValue := ULQ.PositionAddressInfoQuery.QueryRow(
		"insert into PositionAddressInfo (userCode, locLatitude, locLongtitue) values (%s,%s,%s)",
		PositionAddressInfo.UserCode,
		PositionAddressInfo.LocLatitude,
		PositionAddressInfo.LocLongtitue)
	if resValue.Err() != nil {
		return log.MyError(resValue.Err().Error())
	}
	return nil
}

func (ULQ *ULQuery) PositionAddressInfoRead(ul models.PositionAddressInfo) ([]models.PositionAddressInfo, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	return nil, nil
}

func (ULQ *ULQuery) PositionAddressInfoUpdate(ul models.PositionAddressInfo) (bool, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	return false, nil
}

func (ULQ *ULQuery) PositionAddressInfoDelete(ul models.PositionAddressInfo) (bool, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	return false, nil
}
