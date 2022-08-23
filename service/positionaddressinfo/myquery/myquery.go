package myquery

import (
	"database/sql"
	"fmt"

	"github.com/jeonyunjae/fiber-api/datatype/query"
	"github.com/jeonyunjae/fiber-api/models"
	"github.com/jeonyunjae/fiber-api/util/log"
)

var PositionAddressInfo ULQuery

type ULQuery struct {
	PositionAddressInfoQuery query.DBInstance
}

func (ULQ *ULQuery) PositionAddressInfoInit() error {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	if PositionAddressInfo.PositionAddressInfoQuery.Db == nil {
		return log.MyError("Error_PositionAddressInfoInit")
	}
	PositionAddressInfo.PositionAddressInfoQuery = query.Database

	return nil
}

func (ULQ *ULQuery) PositionAddressInfoInsert(ul models.Positionaddressinfo) error {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	var sql = fmt.Sprintf(`INSERT INTO public.PositionAddressInfos(
	userCode, 
	locLatitude, 
	locLongtitude
	)VALUES ('%s', %f, %f);`, ul.Usercode, ul.Loclatitude, ul.Loclongtitude)

	err := ULQ.PositionAddressInfoQuery.Insert(sql)

	if err != nil {
		return log.MyError(err.Error())
	}

	// data, err := ULQ.PositionAddressInfoRead(ul)
	// if len(data) < 1 || err != nil {
	// 	return err
	// }

	return nil
}

func (ULQ *ULQuery) PositionAddressInfoRead(ul models.Positionaddressinfo) (map[string]models.Positionaddressinfo, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	//m := make(map[string]models.Positionaddressinfo)

	var sql = fmt.Sprintf(`SELECT userCode, locLatitude, locLongtitude FROM public.PositionAddressInfos where userCode = '%s'`, ul.Usercode)

	rows, err := ULQ.PositionAddressInfoQuery.Select(sql)
	if err != nil {
		return nil, err
	}
	m, err := dataScan(rows)
	if err != nil {
		return m, log.MyError("Error_PositionAddressInfoAllRead")
	}

	return nil, nil
}

func (ULQ *ULQuery) PositionAddressInfoAllRead() (map[string]models.Positionaddressinfo, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	var sql = `SELECT userCode, locLatitude, locLongtitude FROM public.PositionAddressInfos`

	rows, err := ULQ.PositionAddressInfoQuery.Select(sql)

	if err != nil {
		return nil, err
	}

	m, err := dataScan(rows)
	if err != nil {
		return m, log.MyError("Error_PositionAddressInfoAllRead")
	}
	return m, nil
}

func dataScan(rows *sql.Rows) (map[string]models.Positionaddressinfo, error) {
	var userCode string
	var locLatitude, locLongtitude float64

	m := make(map[string]models.Positionaddressinfo)

	for rows.Next() {
		err := rows.Scan(&userCode, &locLatitude, &locLongtitude)
		if err != nil {
			return m, log.MyError("Error_dataScan")
		}
		m[userCode] = models.Positionaddressinfo{Usercode: userCode, Loclatitude: locLatitude, Loclongtitude: locLongtitude}
	}
	return m, nil
}

func (ULQ *ULQuery) PositionAddressInfoAllReads(args ...any) (map[string]models.Positionaddressinfo, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	var userCode string
	var locLatitude, locLongtitude float64

	m := make(map[string]models.Positionaddressinfo)
	var sql string = `SELECT "userCode",earth_distance(ll_to_earth("locLatitude","locLongtitude"),ll_to_earth(37.482325, 126.881754))
		FROM public."PositionAddressInfo" ORDER BY earth_distance DESC`

	rows, err := ULQ.PositionAddressInfoQuery.Select(sql)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&userCode, &locLatitude, &locLongtitude)
		if err != nil {
			return m, log.MyError("Error_PositionAddressInfoAllReads")
		}
		m[userCode] = models.Positionaddressinfo{Usercode: userCode, Loclatitude: locLatitude, Loclongtitude: locLongtitude}
	}

	return m, nil
}

func (ULQ *ULQuery) PositionAddressInfoUpdate(ul models.Positionaddressinfo) (bool, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	return false, nil
}

func (ULQ *ULQuery) PositionAddressInfoDelete(ul models.Positionaddressinfo) (bool, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	return false, nil
}
