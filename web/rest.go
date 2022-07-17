package web

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"persons-web-micro/db"
	"persons-web-micro/models"
	"strconv"
)

func Cats(rout *mux.Router) {
	rout.HandleFunc("/rest/api/cats", func(wr http.ResponseWriter, req *http.Request) {
		var cats []models.Cat
		cats = db.GetCats(models.Cat{})
		jsonCats, err := json.Marshal(cats)
		if err != nil {
			return
		}
		_, err = fmt.Fprintln(wr, string(jsonCats))
		if err != nil {
			return
		}
	})
}

func CatById(rout *mux.Router) {
	rout.HandleFunc("/rest/api/cats/{id}", func(wr http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		idS := vars["id"]
		id, err := strconv.Atoi(idS)
		if err != nil {
			return
		}
		cat := db.GetCat(id)
		catJson, err := json.Marshal(cat)
		if err != nil {
			return
		}
		_, err2 := fmt.Fprint(wr, catJson)
		if err2 != nil {
			return
		}
	})
}

func CountriesRest(rout *mux.Router) {
	rout.HandleFunc("/rest/api/country", func(wr http.ResponseWriter, req *http.Request) {
		countries := db.GetCountries()
		jsonCountr, err := json.Marshal(countries)
		if err != nil {
			return
		}
		_, err = fmt.Fprint(wr, string(jsonCountr))
		if err != nil {
			return
		}
	})
}
