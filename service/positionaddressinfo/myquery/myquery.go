package myquery

import (
	"database/sql"
	"fmt"

	"github.com/jeonyunjae/fiber-api/database/mydbquery"
	"github.com/jeonyunjae/fiber-api/models"
	"github.com/jeonyunjae/fiber-api/util/log"
)

var PositionAddressInfo ULQuery

type ULQuery struct {
	PositionAddressInfoQuery *sql.DB
}

func (ULQ *ULQuery) PositionAddressInfosInit() error {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	if mydbquery.Database.Db == nil {
		return log.MyError("Error_PositionAddressInfosInit")
	}
	PositionAddressInfo.PositionAddressInfoQuery = mydbquery.Database.Db

	return nil
}

func (ULQ *ULQuery) PositionAddressInfoInsert(ul models.PositionAddressInfo) error {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	var sql = fmt.Sprintf(`INSERT INTO public."PositionAddressInfo"(
	"userCode", 
	"locLatitude", 
	"locLongtitude"
	)VALUES ('%s', %f, %f);`, ul.UserCode, ul.LocLatitude, ul.LocLongtitude)

	resValue := ULQ.PositionAddressInfoQuery.QueryRow(
		sql)
	if resValue.Err() != nil {
		return log.MyError(resValue.Err().Error())
	}
	return nil
}

func (ULQ *ULQuery) PositionAddressInfoRead(ul models.PositionAddressInfo) ([]models.PositionAddressInfo, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	return nil, nil
}

func (ULQ *ULQuery) PositionAddressInfoAllRead() (map[string]models.PositionAddressInfo, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	var userCode string
	var locLatitude, locLongtitude float64

	m := make(map[string]models.PositionAddressInfo)

	rows, err := ULQ.PositionAddressInfoQuery.Query(`SELECT "userCode", "locLatitude", "locLongtitude" FROM public."PositionAddressInfo"`)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&userCode, &locLatitude, &locLongtitude)
		if err != nil {
			return m, log.MyError("Error_PositionAddressInfoAllRead")
		}
		m[userCode] = models.PositionAddressInfo{UserCode: userCode, LocLatitude: locLatitude, LocLongtitude: locLongtitude}
	}

	return m, nil
}

func (ULQ *ULQuery) PositionAddressInfoAllReads(args ...any) (map[string]models.PositionAddressInfo, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	var userCode string
	var locLatitude, locLongtitude float64

	m := make(map[string]models.PositionAddressInfo)
	var sql string = `SELECT "userCode",earth_distance(ll_to_earth("locLatitude","locLongtitude"),ll_to_earth(37.482325, 126.881754))
		FROM public."PositionAddressInfo" ORDER BY earth_distance DESC`

	rows, err := ULQ.PositionAddressInfoQuery.Query(sql)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&userCode, &locLatitude, &locLongtitude)
		if err != nil {
			return m, log.MyError("Error_PositionAddressInfoAllReads")
		}
		m[userCode] = models.PositionAddressInfo{UserCode: userCode, LocLatitude: locLatitude, LocLongtitude: locLongtitude}
	}

	return m, nil
}

func (ULQ *ULQuery) PositionAddressInfoUpdate(ul models.PositionAddressInfo) (bool, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	return false, nil
}

func (ULQ *ULQuery) PositionAddressInfoDelete(ul models.PositionAddressInfo) (bool, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	return false, nil
}
