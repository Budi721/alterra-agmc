package seeder

import (
	"github.com/Budi721/alterra-agmc/v6/database"
	"gorm.io/gorm"
)

type seed struct {
	DB *gorm.DB
}

func NewSeeder() *seed {
	return &seed{database.GetConnection()}
}

func (s *seed) SeedAll() {
	userSeeder(s.DB)
}

func (s *seed) DeleteAll() {
	s.DB.Exec("DELETE FROM users")
}
