package models

type Food struct {
	Image    string `xorm:"image"`
	Name     string
	ID       string
	Price    int
	Describe string
}
