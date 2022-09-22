package database

import (
	"fmt"
	"github.com/Budi721/alterra-agmc/v6/pkg/util"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	dbConfig struct {
		Host string
		User string
		Pass string
		Port string
		Name string
	}

	mysqlConfig struct {
		dbConfig
	}
)

func (conf mysqlConfig) Connect() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		conf.Host,
		conf.User,
		conf.Pass,
		conf.Name,
		conf.Port,
	)

	dsn = util.Getenv("DATABASE_URL", dsn)

	var err error

	dbConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
