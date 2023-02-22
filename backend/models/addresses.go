package models

type district struct {
	Districts string
	Area      string
	City      string
	Province  string
}

type Address struct {
	Id           int
	AcceptName   string
	Mobile       string
	ProvinceName string
	Area         int
	City         int
	Sex          uint
	Street       string
	Inner        bool
	Lat          string
	DoorNumber   string
	AreaName     string
	CityName     string
	PoiName      string
	District     []district
}
