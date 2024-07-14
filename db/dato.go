package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DNS = "host=localhost user=wander password=test12 dbname=arduino"
var DB *gorm.DB

func Conex() {
	var err error
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		log.Println("fallo en la conexcion...")
	}
	log.Println("DBconectado...")
}
