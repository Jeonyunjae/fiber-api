package list

import (
	"fmt"

	"github.com/jeonyunjae/fiber-api/models"
	"github.com/jeonyunjae/fiber-api/service/userlocation"
)

var UserlocationList ULList

type ULList struct {
	UserlocationList []models.UserLocation
}

func (ULL *ULList) UserlocationsInit() ULList {
	fmt.Println("UserlocationListInit")
	ULL.UserlocationList = userlocation.UserLocationCsvToStruct()
	return *ULL
}

func (ULL *ULList) CreateUserlocationList(ul models.UserLocation) ULList {
	fmt.Println("UserlocationListCreate")
	ULL.UserlocationList = append(ULL.UserlocationList, ul)
	return *ULL
}

func (ULL *ULList) ReadUserlocationList(ul models.UserLocation) ULList {
	fmt.Println("UserlocationListRead")
	return *ULL
}

func (ULL *ULList) UpdateUserlocationList(ul models.UserLocation) models.UserLocation {
	fmt.Println("UserlocationListUpdate")
	var row models.UserLocation
	for _, row = range ULL.UserlocationList {
		if row.ID == ul.ID {
			row.Lat = ul.Lat
			row.Lon = ul.Lon
			row.CityCode = ul.CityCode
			return row
		}
	}
	return row
}

func (ULL *ULList) DeleteUserlocationList(ul models.UserLocation) ULList {
	fmt.Println("UserlocationListDelete")
	ULL.UserlocationList = userlocation.UserLocationCsvToStruct()
	return *ULL
}
