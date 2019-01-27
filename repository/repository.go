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

func EmailIsUsed(db *gorm.DB, email string) bool {
	u := new(models.User)
	return !db.Where("email = ?", email).First(u).RecordNotFound()
}

func CreateUser(db *gorm.DB, u *models.User) error {
	return db.Save(u).Error
}


func GetUsersByName(db *gorm.DB, name string) ([]*models.User, error) {
	users := make([]*models.User, 0)
	err := db.Where("name LIKE ?", "%"+name+"%").Limit(10).Find(&users).Error
	return users, err
}




/*
func GetLastKey(db *gorm.DB) (*models.SecretKey, error) {
	k := new(models.SecretKey)
	err := db.Last(k).Error
	return k, err
}

func UpdateKey(db *gorm.DB, k *models.SecretKey) error {
	return db.Save(k).Error
}

func GetKey(db *gorm.DB, key string) (*models.SecretKey, error) {
	k := new(models.SecretKey)
	err := db.Where("Key = ?", key).First(k).Error
	return k, err
}

func FreeKeyCount(db *gorm.DB) (int, error) {
	var count int
	err := db.Model(models.SecretKey{}).Where("sent_at is NULL").Count(&count).Error
	return count, err
}
*/