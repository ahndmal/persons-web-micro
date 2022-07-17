package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"persons-web-micro/models"
)

func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("WELT_IP"),
		os.Getenv("WELT_USER"),
		os.Getenv("WELT_PASS"),
		os.Getenv("WELT_DB"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetCats(model interface{}) []models.Cat {
	model = &models.Cat{}
	db, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&model)

	cats := make([]models.Cat, 0)

	db.Find(&cats) // todo - change

	return cats
}

func GetCat(id int) models.Cat {
	var cat models.Cat
	db, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&cat)
	db.Find(cat, "id = ?", id)
	return cat
}
