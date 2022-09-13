package database

import (
	"github.com/Budi721/alterra-agmc/v2/config"
	"github.com/Budi721/alterra-agmc/v2/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}
