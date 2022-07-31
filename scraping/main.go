package main

import (
	. "housing_viz/common"
	"housing_viz/scraping/scrapers"
	"time"
)

func main() {
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
	}.start()
}
