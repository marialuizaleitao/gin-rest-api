package database

import (
	"github.com/marialuizaleitao/gin-rest-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic("Failed to connect to database")
	}
	DB.AutoMigrate(&models.Member{})
}
