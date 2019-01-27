package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(256);unique_index"`
	Name     string `gorm:"type:varchar(256);index"`
	Password string `json:"-"`
	Balance  uint
	Token    string `gorm:"-"`
}

type Transaction struct {
	gorm.Model
	FromID   uint `gorm:"index"`
	ToID     uint `gorm:"index"`
	Amount   uint
}
