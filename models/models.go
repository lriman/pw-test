package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email";gorm:"unique_index"`
	Name     string `json:"name";gorm:"index"`
	Password string `json:"-"`
	Balance  uint   `json:"balance"`
	Token    string `json:"token";gorm:"-"`
}

type Transactions struct {
	gorm.Model
	FromID   uint `gorm:"index"`
	ToID     uint `gorm:"index"`
	Amount   uint
}
