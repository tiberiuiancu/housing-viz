package web

import (
	"go.mongodb.org/mongo-driver/bson"
	"housing_viz/pkg/common"
	"html/template"
	"log"
	"net/http"
	"os"
)

var db common.MongoConn

type mapInfo struct {
	MapPoints []mapPoint
	CenterLat float64
	CenterLng float64
	ApiKey    string
}

func mapHandler(w http.ResponseWriter, r *http.Request) {

	res, _ := db.FindAllMaxN(
		bson.D{{"city", "Amsterdam"}},
		-1,
	)

	info := mapInfo{
		convertListingsToMapPoint(res),
		52.358873,
		4.861,
		os.Getenv("MAPS_API_KEY"),
	}

	// parse and render template
	t, _ := template.ParseFiles("templates/map.gohtml")
	if err := t.Execute(w, &info); err != nil {
		log.Println(err)
	}
}

func InitServer() {
	// init db connection
	if err := db.InitConn(); err != nil {
		log.Fatal(err)
	}

	// add handlers
	http.HandleFunc("/", mapHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
