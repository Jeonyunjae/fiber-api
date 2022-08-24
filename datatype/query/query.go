package query

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/BurntSushi/toml"
	"github.com/jeonyunjae/fiber-api/util/config"
	"github.com/jeonyunjae/fiber-api/util/log"
)

type DBInstance struct {
	Db *sql.DB
}

var Database DBInstance

func ConnectDb() {
	var dbStruct config.DbDriverStruct
	_, err := toml.DecodeFile("config.toml", &dbStruct)
	if err == nil {

		psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			dbStruct.Entity.Host,
			dbStruct.Entity.Port,
			dbStruct.Entity.User,
			dbStruct.Entity.Password,
			dbStruct.Entity.DbName)

		// open database
		db, err := sql.Open("postgres", psqlconn)

		if err != nil {
			log.MyError(err.Error())
		}
		Database = DBInstance{Db: db}

		//defer Database.close()
		//Database.ping() //만약 db연결이 실패할 경우 왜 실패한건지 에러찍는 용도
		if err != nil || db.Ping() != nil {
			panic(err.Error())
		}
	}
}

func (DBInstance *DBInstance) close() {
	//DBInstance.Db.Close()
}

func (DBInstance *DBInstance) Select(sql string) (*sql.Rows, error) {
	defer Database.close()

	rows, err := DBInstance.Db.Query(sql)
	if rows == nil || err != nil {
		return nil, log.MyError(rows.Err().Error())
	}
	return rows, err
}

func (DBInstance *DBInstance) Insert(sql string) error {
	defer Database.close()

	rows, err := DBInstance.Db.Query(sql)
	if rows == nil || err != nil {
		return log.MyError(rows.Err().Error())
	}
	return err
}

func (DBInstance *DBInstance) Update(sql string) error {
	defer Database.close()

	rows, err := DBInstance.Db.Query(sql)
	if rows == nil || err != nil {
		return log.MyError(rows.Err().Error())
	}
	return err
}

func (DBInstance *DBInstance) Delete(sql string) error {
	defer Database.close()

	rows, err := DBInstance.Db.Query(sql)
	if rows == nil || err != nil {
		return log.MyError(rows.Err().Error())
	}
	return err
}
