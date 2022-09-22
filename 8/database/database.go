package database

import (
	"sync"

	"github.com/Budi721/alterra-agmc/v6/pkg/util"
	"gorm.io/gorm"
)

var (
	dbConn *gorm.DB
	once   sync.Once
)

func CreateConnection() {
	conf := dbConfig{
		User: util.Getenv("DB_USER", "postgres"),
		Pass: util.Getenv("DB_PASS", "root"),
		Host: util.Getenv("DB_HOST", "database"),
		Port: util.Getenv("DB_PORT", "5432"),
		Name: util.Getenv("DB_NAME", "training"),
	}

	mysql := mysqlConfig{dbConfig: conf}
	once.Do(func() {
		mysql.Connect()
	})
}

func GetConnection() *gorm.DB {
	if dbConn == nil {
		CreateConnection()
	}
	return dbConn
}
