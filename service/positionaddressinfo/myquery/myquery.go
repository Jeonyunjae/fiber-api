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

	if query.Database.Db == nil {
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
	name,
	address
	)VALUES ('%s', %f, %f,'%s','%s');`, ul.Usercode, ul.Loclatitude, ul.Loclongtitude, ul.Name, ul.Address)

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

	var sql = fmt.Sprintf(`SELECT userCode, locLatitude, locLongtitude, name, address FROM public.PositionAddressInfos where userCode = '%s'`, ul.Usercode)

	rows, err := ULQ.PositionAddressInfoQuery.Select(sql)
	if err != nil {
		return nil, err
	}
	m, err := dataScanByPositionAddressInfo(rows)
	if err != nil {
		return m, log.MyError("Error_PositionAddressInfoAllRead")
	}

	return nil, nil
}

func (ULQ *ULQuery) PositionAddressInfoReads(ul models.PositionaddressDistanceInfo) ([]models.PositionaddressDistanceInfo, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	//m := make(map[string]models.Positionaddressinfo)

	var sql = fmt.Sprintf(`select usercode, loclatitude, loclongtitude, name, address,earth_distance(ll_to_earth(loclatitude, loclongtitude), ll_to_earth(%f, %f)) as distance from public.positionaddressinfos order by distance asc limit %d`, ul.Loclatitude, ul.Loclongtitude, ul.Count)

	rows, err := ULQ.PositionAddressInfoQuery.Select(sql)
	if err != nil {
		return nil, err
	}
	m, err := dataScanByPositionAddressDistanceInfo(rows)
	if err != nil {
		return m, log.MyError("Error_PositionAddressInfoAllRead")
	}

	return m, nil
}

func (ULQ *ULQuery) PositionAddressInfoReadsRange(ul models.PositionaddressDistanceInfo) ([]models.PositionaddressDistanceInfo, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()
	//m := make(map[string]models.Positionaddressinfo)

	var sql = fmt.Sprintf(`select usercode, loclatitude,loclongtitude , name, address, earth_distance(ll_to_earth(loclatitude, loclongtitude), ll_to_earth(%f, %f)) as distance from public.positionaddressinfos order by distance asc limit %d`, ul.Loclatitude, ul.Loclongtitude, ul.Count)

	rows, err := ULQ.PositionAddressInfoQuery.Select(sql)
	if err != nil {
		return nil, err
	}
	m, err := dataScanByPositionAddressDistanceInfo(rows)
	if err != nil {
		return m, log.MyError("Error_PositionAddressInfoAllRead")
	}

	return m, nil
}

func (ULQ *ULQuery) PositionAddressInfoAllRead() (map[string]models.Positionaddressinfo, error) {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	var sql = `SELECT userCode, locLatitude,  name, address,locLongtitude FROM public.PositionAddressInfos`

	rows, err := ULQ.PositionAddressInfoQuery.Select(sql)

	if err != nil {
		return nil, err
	}

	m, err := dataScanByPositionAddressInfo(rows)
	if err != nil {
		return m, log.MyError("Error_PositionAddressInfoAllRead")
	}
	return m, nil
}

func (ULQ *ULQuery) PositionAddressInfoUpdate(ul models.Positionaddressinfo) error {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	var sql = fmt.Sprintf(`UPDATE public.positionaddressinfos SET loclatitude=%f, loclongtitude=%f WHERE usercode='%s'`, ul.Loclatitude, ul.Loclongtitude, ul.Usercode)

	err := ULQ.PositionAddressInfoQuery.Update(sql)

	if err != nil {
		return err
	}

	return nil
}

func (ULQ *ULQuery) PositionAddressInfoDelete(ul models.Positionaddressinfo) error {
	defer log.ElapsedTime(log.TraceFn(), "start")()

	var sql = fmt.Sprintf(`DELETE FROM public.positionaddressinfos WHERE usercode ='%s';`, ul.Usercode)

	err := ULQ.PositionAddressInfoQuery.Delete(sql)

	if err != nil {
		return err
	}

	return nil
}

func dataScanByPositionAddressDistanceInfo(rows *sql.Rows) ([]models.PositionaddressDistanceInfo, error) {
	var userCode, name, address string
	var locLatitude, locLongtitude, distance float64

	var m []models.PositionaddressDistanceInfo

	for rows.Next() {
		err := rows.Scan(&userCode, &locLatitude, &locLongtitude, &name, &address, &distance)
		if err != nil {
			return m, log.MyError("Error_dataScan")
		}
		m = append(m, models.PositionaddressDistanceInfo{Usercode: userCode, Loclatitude: locLatitude, Loclongtitude: locLongtitude, Name: name, Address: address, Distance: distance})
	}
	return m, nil
}

func dataScanByPositionAddressInfo(rows *sql.Rows) (map[string]models.Positionaddressinfo, error) {
	var userCode, name, address string
	var locLatitude, locLongtitude float64

	m := make(map[string]models.Positionaddressinfo)

	for rows.Next() {
		err := rows.Scan(&userCode, &locLatitude, &name, &address, &locLongtitude)
		if err != nil {
			return m, log.MyError("Error_dataScan")
		}
		m[userCode] = models.Positionaddressinfo{Usercode: userCode, Loclatitude: locLatitude, Loclongtitude: locLongtitude, Name: name, Address: address}
	}
	return m, nil
}
