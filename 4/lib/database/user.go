package database

import (
	"github.com/Budi721/alterra-agmc/v2/config"
	"github.com/Budi721/alterra-agmc/v2/middlewares"
	"github.com/Budi721/alterra-agmc/v2/models"
	"gorm.io/gorm"
)

func LoginUser(email string, password string) (string, error) {
	var user models.User

	if err := config.DB.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return "", err
	}

	token, err := middlewares.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
func GetUsers() ([]models.User, error) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return users, nil
}

func GetUser(id uint) (*models.User, error) {
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return &models.User{}, err
	}
	return &user, nil
}

func CreateUser(name string, email string, password string) (*models.User, error) {
	user := models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return &models.User{}, err
	}

	return &user, nil
}

func UpdateUser(id uint, name string, email string, password string) (*models.User, error) {
	user := models.User{
		ID: id,
	}
	if err := config.DB.First(&user).Error; err != nil {
		return &models.User{}, err
	}

	user.Name = name
	user.Email = email
	user.Password = password

	if err := config.DB.Save(&user).Error; err != nil {
		return &models.User{}, err
	}

	return &user, nil
}

func DeleteUser(id uint) (*models.User, error) {
	user := models.User{
		ID: id,
	}

	if err := config.DB.First(&user).Error; err != nil {
		return &models.User{}, err
	}

	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		return &models.User{}, err
	}

	return &user, nil
}
