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

/*
В целом внешнего ключа достаточно,
но если дублировать в логах поля, по которым нужно сделать фильтр,
то это поможет избежать лишних JOIN'ов при запроса
*/
type Transaction struct {
	gorm.Model
	FromID      uint   `gorm:"index"`
	ToID        uint   `gorm:"index"`
	FromName    string `gorm:"index"`
	ToName      string `gorm:"index"`
	Amount      uint   `gorm:"index"`
	FromBalance uint
	ToBalance   uint
}
