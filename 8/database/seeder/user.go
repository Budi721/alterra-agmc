package seeder

import (
	"log"

	"github.com/Budi721/alterra-agmc/v6/internal/model"
	"gorm.io/gorm"
)

func userSeeder(db *gorm.DB) {
	var roles = []model.User{
		{ID: 1, Name: "Budi Rahmawan", Email: "rahmawanbd@gmail.com", Password: "123456"},
		{ID: 2, Name: "Tresno Asih", Email: "tresnodev@golek.com", Password: "qwerty"},
	}
	if err := db.Create(&roles).Error; err != nil {
		log.Printf("cannot seed data roles, with error %v\n", err)
	}
	log.Println("success seed data roles")
}
