package main

import (
	"github.com/pw-test/repository"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"github.com/pw-test/models"
	"regexp"
)

func (a *App) SignInController(email string, password string) (*models.User, error) {

	u, err := repository.GetUserByEmail(a.DB, email)
	if err != nil{
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, fmt.Errorf("login and password does not match")
	}

	u.Token = getToken(email)
	return u, nil
}


func (a *App) SignUpController(email string, pw1 string, pw2 string, name string) error {

	if pw1 != pw2 {
		return fmt.Errorf("password and confirmation is not equal")
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(pw1), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.MatchString(email) {
		return fmt.Errorf("valid email address is required")
	}

	if len(name) < 2 {
		return fmt.Errorf("user name is required")
	}

	isUsed := repository.EmailIsUsed(a.DB, email)
	if isUsed{
		return fmt.Errorf("email already used")
	}

	u := models.User{
		Name: name,
		Email: email,
		Password: string(passwordHash),
		Balance: 500,
	}

	return repository.SaveUser(a.DB, &u)
}


func (a *App) ProfileController(email string) (*models.ResponseProfile, error) {
	u, err := repository.GetUserByEmail(a.DB, email)
	if err != nil{
		return nil, err
	}

	return &models.ResponseProfile{ID: u.ID, Name: u.Name, Balance: u.Balance, Email: u.Email}, nil
}

func (a *App) AutoCompleteController(name string) ([]*models.ResponseAutoComplete, error) {

	users, err := repository.GetUsersByName(a.DB, name)
	if err != nil{
		return nil, err
	}

	profiles := make([]*models.ResponseAutoComplete, 0)
	for _, u := range users{
		profiles = append(profiles, &models.ResponseAutoComplete{ID: u.ID, Name: u.Name, Email: u.Email})
	}

	return profiles, nil
}

func (a *App) TransferController(from string, to string, amount uint) error {

	var err error
	var fromUser, toUser *models.User

	tx := a.DB.Begin()
	if err = tx.Error; err != nil {
		return err
	}

	fromUser, err = repository.GetUsersForUpdate(a.DB, from)
	if err != nil{
		tx.Rollback()
		return err
	}

	toUser, err = repository.GetUsersForUpdate(a.DB, to)
	if err != nil{
		tx.Rollback()
		return err
	}

	if fromUser.Balance < amount {
		tx.Rollback()
		return fmt.Errorf("balance is too low")
	}

	//txLog := models.Transaction{FromID: fromUser.ID, ToID: toUser.ID, Amount: amount}
	fromUser.Balance -= amount
	toUser.Balance += amount

	err = tx.Save(fromUser).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Save(toUser).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Save(&models.Transaction{FromID: fromUser.ID, ToID: toUser.ID, Amount: amount}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
