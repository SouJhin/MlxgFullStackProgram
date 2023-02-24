package models

import "time"

type AttendanceList struct {
	AttendanceList   int
	UpdatedAt        time.Time `xorm:"updated"`
	Id               int
	CreatedAt        time.Time `xorm:"created"`
	AttendanceCoupon int
	Nickname         string
	Num              int
	RewardDays       int
	UserId           string
	Date             time.Time
	AttendancePoint  int
	StoreId          int
}
