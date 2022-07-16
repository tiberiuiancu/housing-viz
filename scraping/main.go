package main

import "time"

func main() {
	Scheduler{
		scrapers: []Scraper{
			{
				name:        "Dummy",
				exec:        dummyScraperRun,
				channel:     make(chan *Listing),
				nextRunTime: time.Now(),
				cooldown:    time.Second * 10,
				isRunning:   false,
			},
		},
	}.start()
}
