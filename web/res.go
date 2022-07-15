package web

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"persons-web-micro/db"
	"strconv"
)

func Home() {
	http.HandleFunc("/", func(wr http.ResponseWriter, request *http.Request) {
		meth := request.Method
		fmt.Println(meth)

		_, err := fmt.Fprint(wr, "HOME Page")
		if err != nil {
			return
		}
	})
}

func Cats() {
	http.HandleFunc("/cats", func(wr http.ResponseWriter, req *http.Request) {
		var cats []db.Cat
		cats = db.GetCats(db.Cat{})
		_, err := fmt.Fprintln(wr, cats)
		if err != nil {
			return
		}
	})
}

func CatById() {
	http.HandleFunc("/cats/{id}", func(wr http.ResponseWriter, req *http.Request) {
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
