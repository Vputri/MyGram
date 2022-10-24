package database

import (
	"fmt"
	"log"
	"middleware/models"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var (
	host     = "containers-us-west-105.railway.app"
	user     = "postgres"
	password = "7P7Ia9JzEPlWJ9fwHoYl"
	dbPort   = "5570"
	dbname   = "railway"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	fmt.Println("sukses koneksi ke database")
	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
}

func GetDB() *gorm.DB {
	return db
}
