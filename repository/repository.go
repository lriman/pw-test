package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/pw-test/models"
)

func GetUserByEmail(db *gorm.DB, email string) (*models.User, error) {
	u := new(models.User)
	err := db.Where("email = ?", email).First(u).Error
	return u, err
}

func GetUsersForUpdate(db *gorm.DB, email string) (*models.User, error) {
	u := new(models.User)
	err := db.Where("email = ?", email).Set("gorm:query_option", "FOR UPDATE").First(u).Error
	return u, err
}

func EmailIsUsed(db *gorm.DB, email string) bool {
	u := new(models.User)
	return !db.Where("email = ?", email).First(u).RecordNotFound()
}

func SaveUser(db *gorm.DB, u *models.User) error {
	return db.Save(u).Error
}

func GetUsersByName(db *gorm.DB, name string) ([]*models.User, error) {
	users := make([]*models.User, 0)
	err := db.Where("name LIKE ?", "%"+name+"%").Limit(10).Find(&users).Error
	return users, err
}
