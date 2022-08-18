package main

import (
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
