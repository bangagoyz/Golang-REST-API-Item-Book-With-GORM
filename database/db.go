package database

import (
	// "chapter2_4/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "naha22"
	dbPort   = "5432"
	dbname   = "learning-gorm"
	db       *gorm.DB
	err      error
)

func StartDB() (*gorm.DB, error) {
	config := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error menghubungkan ke database: ", err)
	}

	// db.Debug().AutoMigrate(models.Book{})
	return db, err
}

func GetDB() *gorm.DB {
	return db
}
