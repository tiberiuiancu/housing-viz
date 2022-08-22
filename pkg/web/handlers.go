package web

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"housing_viz/pkg/common"
	"html/template"
	"log"
	"net/http"
	"os"
)

var db common.MongoConn

type mapInfo struct {
	ApiKey string
}

type dataRequest struct {
	LatNE float64 `json:"lat_ne"`
	LngNE float64 `json:"lng_ne"`
	LatSW float64 `json:"lat_sw"`
	LngSW float64 `json:"lng_sw"`
}

func mapHandler(w http.ResponseWriter, r *http.Request) {

	// parse and render template
	t, _ := template.ParseFiles("templates/map.gohtml")
	if err := t.Execute(w, &mapInfo{
		os.Getenv("MAPS_API_KEY"),
	}); err != nil {
		log.Println(err)
	}
}

func getDataHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request", r.Body)
	response := "[]"

	decoder := json.NewDecoder(r.Body)
	var req dataRequest
	if err := decoder.Decode(&req); err != nil {
		log.Println("Failed to decode", r.Body)
	} else {
		// find listings within map area
		if queryResult, err := db.FindAll(bson.D{
			{"lat", bson.D{{"$gte", req.LatSW}}},
			{"lat", bson.D{{"$lte", req.LatNE}}},
			{"lng", bson.D{{"$gte", req.LngSW}}},
			{"lng", bson.D{{"$lte", req.LngNE}}},
		}); err != nil {
			log.Println("Error while fetching data:", err)
		} else {
			mapPoints := convertListingsToMapPoint(queryResult)
			if res, err := json.Marshal(mapPoints); err != nil {
				log.Println("Unable to marshal query result:", err)
			} else {
				response = string(res)
			}
		}
	}

	if _, writeErr := fmt.Fprint(w, response); writeErr != nil {
		log.Println("Unable to send response:", writeErr)
	}
}

func InitServer() {
	// init db connection
	if err := db.InitConn(); err != nil {
		log.Fatal(err)
	}

	// add handlers
	http.HandleFunc("/", mapHandler)
	http.HandleFunc("/data", getDataHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
