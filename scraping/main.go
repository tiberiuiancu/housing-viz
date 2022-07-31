package main

import (
	. "housing_viz/common"
	"housing_viz/scraping/scrapers"
)

func main() {
	//Scheduler{
	//	scrapers: []Scraper{
	//		{
	//			name:        "Dummy",
	//			exec:        scrapers.DummyScraperRun,
	//			channel:     make(chan *Listing),
	//			nextRunTime: time.Now(),
	//			cooldown:    time.Second * 10,
	//			isRunning:   false,
	//		},
	//	},
	//}.start()

	scrapers.ParariusScraperRun(make(chan *Listing))
}
