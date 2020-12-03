package service

import (
	"errors"

	"Walter0697/GinBackend/utility"
	"Walter0697/GinBackend/base"
	"Walter0697/GinBackend/models"
	"Walter0697/GinBackend/apibody"
)

func CreateAccount(username string, password string, userrole uint) (*models.User) {
	hashedPassword := utility.GetEncrpytPassword(password)
	user := models.User{Username: username, Password: hashedPassword, Userrole: userrole}
	base.DB.Create(&user)
	return &user
}

func ChangePassword(input apibody.UpdateAccountInput) (bool, error) {
	var user models.User
	if err := base.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		return false, err
	}

	if !utility.ComparePassword(user.Password, input.OldPassword) {
		return false, errors.New("incorrect password")
	}

	hashedPassword := utility.GetEncrpytPassword(input.NewPassword)

	base.DB.Model(&user).Updates(map[string]interface{}{"password": hashedPassword})

	return true, nil
}

func FindUserByName(username string)  (*models.User, error) {
	var user models.User
	if err := base.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func FindAllUsers() []models.User {
	var users []models.User
	base.DB.Find(&users)
	return users
}