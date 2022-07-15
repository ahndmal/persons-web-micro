package web

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"persons-web-micro/db"
	"strconv"
	"text/template"
	"time"
)

//var (
//	rout = mux.NewRouter()
//)

func Home(rout *mux.Router) {

	type data struct {
		Now time.Time
	}

	rout.HandleFunc("/", func(wr http.ResponseWriter, request *http.Request) {
		tmpl := template.Must(template.ParseFiles("files/home.html"))
		now := time.Now()
		data := data{now}
		err := tmpl.Execute(wr, data)
		if err != nil {
			return
		}
	})
}

func Cats(rout *mux.Router) {
	rout.HandleFunc("/cats", func(wr http.ResponseWriter, req *http.Request) {
		var cats []db.Cat
		cats = db.GetCats(db.Cat{})
		_, err := fmt.Fprintln(wr, cats)
		if err != nil {
			return
		}
	})
}

func CatById(rout *mux.Router) {
	rout.HandleFunc("/cats/{id}", func(wr http.ResponseWriter, req *http.Request) {
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
