package main

import (
	"log"
	"my-blog/handler"
	"net/http"
	"text/template"
)

// mainの関数
func main() {
	h := handler.New(
		template.Must(template.ParseFiles("assets/index.html")),
	)
	http.HandleFunc("/", h.Index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
