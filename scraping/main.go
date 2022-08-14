package main

import (
	. "housing_viz/common"
	"housing_viz/scraping/scrapers"
	"log"
	"time"
)

func main() {
	log.Println("Initializing mongodb connection")
	var db MongoConn

	if err := db.InitConn(); err != nil {
		panic(err)
	}

	log.Println("Starting scheduler")
	Scheduler{
		scrapers: []Scraper{
			{
				name:        "Pararius",
				exec:        scrapers.ParariusScraperRun,
				channel:     make(chan *Listing),
				nextRunTime: time.Now(),
				cooldown:    time.Second * 1000,
				isRunning:   false,
			},
		},
	}.start(db)
}
