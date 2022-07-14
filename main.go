package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Cat struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func main() {
	dsn := "host=172.17.0.2 user=dev password=possum dbname=cats port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&Cat{})

	var cat Cat
	cats := make([]Cat, 0)

	db.First(&cat, "name = ?", "Murz")
	db.Find(&cats, "name like ?", "Mu")

	fmt.Println(cat)
	fmt.Println(cats)
}
