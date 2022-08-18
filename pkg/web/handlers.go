package web

import (
	"go.mongodb.org/mongo-driver/bson"
	"housing_viz/pkg/common"
	"html/template"
	"log"
	"net/http"
)

var db common.MongoConn

func homeHandler(w http.ResponseWriter, r *http.Request) {
	res, _ := db.FindAllMaxN(
		bson.D{{}},
		-1,
	)

	t, _ := template.ParseFiles("templates/listing.html")
	t.Execute(w, res)
}

func InitServer() {
	// init db connection
	if err := db.InitConn(); err != nil {
		log.Fatal(err)
	}

	// add handlers
	http.HandleFunc("/", homeHandler)
	log.Fatal("hmm", http.ListenAndServe(":8080", nil))
}
