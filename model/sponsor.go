package model

import "github.com/jinzhu/gorm"

// Sponsor struct
type Sponsor struct {
	gorm.Model
	Name   string `gorm:"not null" json:"name"`
	Link   string `gorm:"not null" json:"link"`
	Amount int    `gorm:"not null" json:"amount"`
}
