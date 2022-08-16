package models

type PostAddressCodeDefine struct {
	CountryCode  uint   `json:"CountryCode"`
	CountryName  string `json:"CountryName"`
	Address1Code uint   `json:"Address1Code"`
	Address1Name string `json:"Address1Name"`
	Address2Code uint   `json:"Address2Code"`
	Address2Name string `json:"Address2Name"`
	Address3Code uint   `json:"Address3Code"`
	Address3Name string `json:"Address3Name"`
	AddressCode  uint   `json:"AddressCode" gorm:"primaryKey"`
}
