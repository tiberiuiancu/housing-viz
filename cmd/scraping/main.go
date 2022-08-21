package main

import (
	"go.mongodb.org/mongo-driver/bson"
	. "housing_viz/pkg/common"
	. "housing_viz/pkg/scraping"
	"housing_viz/pkg/scraping/scrapers"
	"log"
	"os"
	"time"
)

func main() {
	// load env variables if we're not running in docker
	if os.Getenv("DOCKER") != "TRUE" {
		LoadEnv()
	}

	log.Println("Initializing mongodb connection")
	var db MongoConn

	if err := db.InitConn(); err != nil {
		panic(err)
	}

	res, _ := db.FindAll(bson.D{{}})
	log.Println("Number of records:", len(res))

	log.Println("Starting scheduler")
	Scheduler{
		Scrapers: []Scraper{
			{
				Name:        "Pararius",
				Exec:        scrapers.ParariusScraperRun,
				Channel:     make(chan *Listing),
				NextRunTime: time.Now(),
				Cooldown:    time.Second * 1000,
				IsRunning:   false,
			},
		},
	}.Start(db)
}
