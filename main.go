package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	web "persons-web-micro/web"
)

func main() {
	rout := mux.NewRouter()
	web.Home(rout)
	web.Cats(rout)
	web.CatById(rout)
	err := http.ListenAndServe(":8082", rout)
	if err != nil {
		log.Fatal(err)
	}
}
