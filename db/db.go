package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type Cat struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func InitDB() (*gorm.DB, error) {
	dsn := os.Getenv("DOCK_PSQL_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetCats(model interface{}) []Cat {
	model = &Cat{}
	db, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&model)

	cats := make([]Cat, 0)

	db.Find(&cats) // todo - change

	return cats
}

func GetCat(id int) Cat {
	var cat Cat
	db, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&cat)
	db.Find(cat, "id = ?", id)
	return cat
}
