package main

import (
	"github.com/joho/godotenv"
	. "housing_viz/common"
	"housing_viz/scraping/scrapers"
	"log"
	"os"
	"time"
)

func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
}

func main() {
	// load env variables if we're not running in docker
	if os.Getenv("DOCKER") == "" {
		loadEnv()
	}

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
