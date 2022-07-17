package web

import (
	"github.com/gorilla/mux"
	"net/http"
	"persons-web-micro/db"
	"persons-web-micro/models"
	"text/template"
	"time"
)

//var (
//	rout = mux.NewRouter()
//)

func Home(rout *mux.Router) {

	type data struct {
		Now  time.Time
		Text string
	}

	rout.HandleFunc("/", func(wr http.ResponseWriter, request *http.Request) {
		tmpl := template.Must(template.ParseFiles("files/home.html"))
		now := time.Now()
		data := data{Now: now, Text: "Hello!"}
		err := tmpl.Execute(wr, data)
		if err != nil {
			return
		}
	})

}

func CatsHtml(rout *mux.Router) {
	type data struct {
		Cats []models.Cat
	}

	rout.HandleFunc("/cats", func(writer http.ResponseWriter, req *http.Request) {
		tmpl := template.Must(template.ParseFiles("files/cats.html"))
		cats := db.GetCats(nil)
		data := data{cats}
		err := tmpl.Execute(writer, data)
		if err != nil {
			return
		}
	})
}
