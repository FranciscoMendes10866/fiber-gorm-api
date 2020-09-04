package model

import "github.com/jinzhu/gorm"

// User struct
type User struct {
	gorm.Model
	Email    string    `gorm:"unique_index;not null" json:"email"`
	Password string    `gorm:"not null" json:"password"`
	Sponsors []Sponsor `gorm:"foreignKey:UserID"`
}

// Sponsor struct
type Sponsor struct {
	gorm.Model
	Name   string `gorm:"not null" json:"name"`
	Link   string `gorm:"not null" json:"link"`
	Amount int    `gorm:"not null" json:"amount"`
	UserID uint
}
