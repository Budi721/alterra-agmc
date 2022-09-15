package config

import (
	"fmt"

	"github.com/Budi721/alterra-agmc/v2/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "root",
		"DB_Port":     "3306",
		"DB_Host":     "database",
		"DB_Name":     "training",
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"],
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitMigrate() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
}
