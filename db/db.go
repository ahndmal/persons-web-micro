package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	}
	db.AutoMigrate(&model)

	var cat Cat
	cats := make([]Cat, 0)

	db.First(&cat, "name = ?", "Murz")
	db.Find(&cats, "name like ?", "Mu")
	//tx := db.Select("select * from cats", "")

	return cats
}
