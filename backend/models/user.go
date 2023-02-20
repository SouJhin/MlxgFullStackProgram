package models

type User struct {
	Id          int `xorm:"not nul"`
	Name        string
	Score       int
	IsStudent   bool
	IsVip       bool
	PhoneNumber string
	Address     string
	Avatar      string
}
