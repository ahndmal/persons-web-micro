package main

import (
	"log"
	"net/http"
	"persons-web-micro/web"
)

func main() {
	web.Home()
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal(err)
	}
}
