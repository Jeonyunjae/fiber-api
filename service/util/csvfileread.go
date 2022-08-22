package util

import (
	"sort"
	"strconv"

	"github.com/jeonyunjae/fiber-api/models"
	"github.com/jeonyunjae/fiber-api/util/excel"
)

func PositionAddressInfoCsvToStruct() []models.Positionaddressinfo {
	rows, _ := excel.FileRead("./fulldata.csv")

	responsePositionAddressInfos := []models.Positionaddressinfo{}

	for _, row := range rows {
		var PositionAddressInfo models.Positionaddressinfo
		PositionAddressInfo.Usercode = row[0]
		PositionAddressInfo.Loclongtitude, _ = strconv.ParseFloat(row[1], 64)
		PositionAddressInfo.Loclatitude, _ = strconv.ParseFloat(row[2], 64)

		responsePositionAddressInfos = append(responsePositionAddressInfos, PositionAddressInfo)
	}
	return responsePositionAddressInfos
}

func PositionAddressInfoCsvToMap() map[string]models.Positionaddressinfo {
	rows, _ := excel.FileRead("./fulldata.csv")

	m := make(map[string]models.Positionaddressinfo)

	for _, row := range rows {
		var PositionAddressInfo models.Positionaddressinfo
		PositionAddressInfo.Usercode = row[0]
		PositionAddressInfo.Loclongtitude, _ = strconv.ParseFloat(row[1], 64)
		PositionAddressInfo.Loclatitude, _ = strconv.ParseFloat(row[2], 64)

		m[PositionAddressInfo.Usercode] = PositionAddressInfo
	}
	return m
}

func PositionAddressInfoCsvToSlice() [][]string {
	rows, _ := excel.FileRead("./fulldata.csv")

	return rows
}

func PostAddressDefineCsvToDb() map[uint]models.PostAddressCodeDefine {
	rows, _ := excel.FileRead("./location.csv")

	m := make(map[uint]models.PostAddressCodeDefine)
	m_temp := make(map[int]string)

	//Address1 추출
	for num, row := range rows {
		if num == 0 {
			continue
		}
		Address1Code, _ := strconv.ParseUint(row[0], 10, 64)

		m_temp[int(Address1Code)] = row[1]
	}
	keys := make([]int, 0, len(m_temp))
	for k := range m_temp {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for num, row := range rows {
		if num == 0 {
			continue
		}
		var postAddressCodeDefine models.PostAddressCodeDefine
		textLen := len(row[0])
		if textLen == 2 {
			postAddressCodeDefine.CountryCode = uint(82)
			postAddressCodeDefine.CountryName = "한국"
			Address1Code, _ := strconv.ParseUint(row[0], 10, 64)
			postAddressCodeDefine.Address1Code = uint(Address1Code)
			postAddressCodeDefine.Address1Name = row[1]
			postAddressCodeDefine.Address2Code = 0
			postAddressCodeDefine.Address2Name = ""
			postAddressCodeDefine.Address3Code = 0
			postAddressCodeDefine.Address3Name = ""
			postAddressCodeDefine.AddressCode = uint(Address1Code * 1000000)

			m[postAddressCodeDefine.AddressCode] = postAddressCodeDefine
		}
	}

	for num, row := range rows {
		if num == 0 {
			continue
		}
		var postAddressCodeDefine models.PostAddressCodeDefine
		textLen := len(row[0])

		if textLen == 5 {

			Address1Code, _ := strconv.ParseUint(row[0][:2], 10, 64)

			postAddressCodeDefine.CountryCode = uint(82)
			postAddressCodeDefine.CountryName = "한국"
			postAddressCodeDefine.Address1Code = m[uint(Address1Code*1000000)].Address1Code
			postAddressCodeDefine.Address1Name = m[uint(Address1Code*1000000)].Address1Name
			Address2Code, _ := strconv.ParseUint(row[0][2:5], 10, 64)
			postAddressCodeDefine.Address2Code = uint(Address2Code)
			postAddressCodeDefine.Address2Name = row[1]
			postAddressCodeDefine.Address3Code = 0
			postAddressCodeDefine.Address3Name = ""
			postAddressCodeDefine.AddressCode = (postAddressCodeDefine.Address1Code * 1000000) + uint(Address2Code*1000)

			m[postAddressCodeDefine.AddressCode] = postAddressCodeDefine
		}
	}

	for num, row := range rows {
		if num == 0 {
			continue
		}
		var postAddressCodeDefine models.PostAddressCodeDefine
		textLen := len(row[0])

		if textLen == 8 {

			Address1Code, _ := strconv.ParseUint(row[0][:2], 10, 64)
			Address2Code, _ := strconv.ParseUint(row[0][2:5], 10, 64)
			Address3Code, _ := strconv.ParseUint(row[0][5:], 10, 64)
			AddressCode, _ := strconv.ParseUint(row[0], 10, 64)

			postAddressCodeDefine.CountryCode = uint(82)
			postAddressCodeDefine.CountryName = "한국"
			postAddressCodeDefine.Address1Code = m[uint(Address1Code*1000000)].Address1Code
			postAddressCodeDefine.Address1Name = m[uint(Address1Code*1000000)].Address1Name
			postAddressCodeDefine.Address2Code = m[uint(Address1Code*1000000+Address2Code*1000)].Address2Code
			postAddressCodeDefine.Address2Name = m[uint(Address1Code*1000000+Address2Code*1000)].Address2Name

			postAddressCodeDefine.Address3Code = uint(Address3Code)
			postAddressCodeDefine.Address3Name = row[1]
			postAddressCodeDefine.AddressCode = uint(AddressCode)

			m[postAddressCodeDefine.AddressCode] = postAddressCodeDefine
		}
	}

	var arrayString [][]string
	for _, row := range m {
		var stringValue []string
		CountryCode := strconv.FormatUint(uint64(row.CountryCode), 10)

		stringValue = append(stringValue, CountryCode)
		stringValue = append(stringValue, row.CountryName)

		Address1Code := strconv.FormatUint(uint64(row.Address1Code), 10)

		stringValue = append(stringValue, Address1Code)
		stringValue = append(stringValue, row.Address1Name)

		Address2Code := strconv.FormatUint(uint64(row.Address2Code), 10)
		if Address2Code != "0" {
			stringValue = append(stringValue, Address2Code)
			stringValue = append(stringValue, row.Address2Name)
		} else {
			stringValue = append(stringValue, Address2Code)
			stringValue = append(stringValue, "_")
		}

		Address3Code := strconv.FormatUint(uint64(row.Address3Code), 10)

		if Address3Code != "0" {
			stringValue = append(stringValue, Address3Code)
			stringValue = append(stringValue, row.Address3Name)
		} else {
			stringValue = append(stringValue, Address3Code)
			stringValue = append(stringValue, "_")
		}
		AddressCode := strconv.FormatUint(uint64(row.AddressCode), 10)
		stringValue = append(stringValue, AddressCode)

		arrayString = append(arrayString, stringValue)
	}

	excel.FileWrite(arrayString)

	return m
}
