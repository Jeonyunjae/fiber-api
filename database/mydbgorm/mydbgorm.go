package mydbgorm

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/jeonyunjae/fiber-api/util/config"
	"github.com/jeonyunjae/fiber-api/util/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	Db *gorm.DB
}

var Database DBInstance

func ConnectDb() {
	var dbStruct config.DbDriverStruct
	_, err := toml.DecodeFile("config.toml", &dbStruct)
	if err == nil {

		psqlconn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
			dbStruct.Entity.Host,
			dbStruct.Entity.User,
			dbStruct.Entity.Password,
			dbStruct.Entity.DbName,
			dbStruct.Entity.Port)

		// open database
		db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})

		if err != nil {
			log.MyError(err.Error())
		}
		log.MyLog("Connected to the GORM successfully")
		db.Logger = logger.Default.LogMode(logger.Info)
		//log.MyLog("Running Migrations")
		//TODO: Add migrations
		//db.AutoMigrate(&models.PositionAddressInfo{})
		Database = DBInstance{Db: db}
	}

}
