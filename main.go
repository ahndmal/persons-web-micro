package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	web "persons-web-micro/web"
)

func main() {
	rout := mux.NewRouter()
	fs := http.FileServer(http.Dir("files/assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	//html
	web.Home(rout)
	web.CatsHtml(rout)
	//rest
	web.Cats(rout)
	web.CatById(rout)
	web.CountriesRest(rout)

	err := http.ListenAndServe(":8082", rout)
	if err != nil {
		log.Fatal(err)
	}
}
