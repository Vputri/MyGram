package database

import (
	"fmt"
	"log"
	"middleware/models"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	dsn := "postgres://kgxrojvgaloxxb:fa9aa0e2416b54f5efb873a2bc06547f429641b64a6a3129091830b38a17933e@ec2-52-3-60-53.compute-1.amazonaws.com:5432/d8qj9cc3do9814"
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
