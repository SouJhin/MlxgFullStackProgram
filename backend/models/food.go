package models

import "time"

type Food struct {
	Image     string    `xorm:"image" json:"image,omitempty"`
	Name      string    `json:"name,omitempty"`
	ID        string    `json:"id,omitempty"`
	Price     int       `json:"price,omitempty"`
	Describe  string    `json:"describe,omitempty"`
	CreatedAt time.Time `xorm:"created" json:"created_at"`
}
