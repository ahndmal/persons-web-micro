package web

import (
	"fmt"
	"net/http"
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
		//
	})
}
