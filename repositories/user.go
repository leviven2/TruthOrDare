package repositories

import (
	"TruthOrDare/database"
	"TruthOrDare/models"
)

func CreateUser(name string, password string) {
	user := models.User{Name: name, Password: password}
	database.DB().Create(&user)
}

func Login(name string, password string) bool {
	var user models.User
	err := database.DB().Where("name = ?", name).First(&user).Error
	if err != nil {
		return false
	}

	return user.Password == password
}

func DoesUserExist(name string) bool {
	var user models.User
	database.DB().Where("name = ?", name).First(&user)
	return user.Name == name
}
