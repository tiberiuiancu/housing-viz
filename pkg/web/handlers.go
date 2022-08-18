package web

import (
	"housing_viz/pkg/common"
	"html/template"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/home.html")
	p := &common.Listing{Url: "123"}
	t.Execute(w, p)
}

func InitServer() {
	// add handlers
	http.HandleFunc("/", homeHandler)
	log.Fatal("hmm", http.ListenAndServe(":8080", nil))
}
