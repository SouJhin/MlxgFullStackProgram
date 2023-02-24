package models

import "time"

type store struct {
	Address   string
	Longitude string
	Latitude  string
	Mobile    string
	Name      string
}

type good struct {
	Number       int
	OriginAmount float64
	Price        float64
	Unit         string
	Property     string
	Image        string
	Amount       float64
	Name         string
}

type Order struct {
	CouponName    string
	ReceiveAt     int
	PayMode       string
	PayUserName   string
	UpdatedAt     string
	GoodsNum      int
	CompletedAt   time.Time
	CreatedAt     time.Time `xorm:"created"`
	InvoiceStatus uint
	SentTime      uint
	StatusText    string
	Remark        string
	CouponAmount  int
	Mobile        string
	UserName      string
	PayedAt       string
	TotalMount    float64
	SendStatus    uint
	Discount      string
	Status        uint
	CompletedTime string
	Amount        float64
	MultiStore    string
	ProductTime   string
	PostScript    string
	SortNum       int
	OrderNo       string
	Id            int
	TypeCate      int
	Goods         []good
	Store         store
}
