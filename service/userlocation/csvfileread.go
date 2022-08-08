package userlocation

import (
	"strconv"

	"github.com/jeonyunjae/fiber-api/models"
	"github.com/jeonyunjae/fiber-api/util/excel"
)

func UserLocationCsvToStruct() []models.UserLocation {
	rows, _ := excel.FileRead("./fulldata.csv")

	responseUserLocations := []models.UserLocation{}

	for _, row := range rows {
		var userlocation models.UserLocation
		userlocation.ID, _ = strconv.ParseUint(row[0], 10, 32)
		userlocation.Lon, _ = strconv.ParseFloat(row[1], 64)
		userlocation.Lat, _ = strconv.ParseFloat(row[2], 64)
		userlocation.CityCode, _ = strconv.ParseUint(row[3], 10, 32)

		responseUserLocations = append(responseUserLocations, userlocation)
	}
	return responseUserLocations
}
