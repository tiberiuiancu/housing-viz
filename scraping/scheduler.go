package main

import (
	. "housing_viz/common"
	"time"
)

type Scheduler struct {
	scrapers []Scraper
}

var syncQueue []Listing

func (s Scheduler) start() {
	for {
		for idx := range s.scrapers {
			scraper := &s.scrapers[idx]

			if scraper.shouldRun() {
				// run if necessary
				scraper.run()
			} else if scraper.isRunning {
				// the goroutine was running; check if the run completed or collect results
				for shouldReceive := true; shouldReceive; {
					select {
					// there is something to receive
					case newListing := <-scraper.channel:
						if newListing != nil {
							// if we receive something other than nil, add it to the sync list
							syncQueue = append(syncQueue, *newListing)
						} else {
							// if we receive nil it's a sign the goroutine finished
							// reschedule and exit for loop
							scraper.reschedule()
							shouldReceive = false
						}
					default:
						// goroutine still running; exit for loop
						shouldReceive = false
					}
				}
			}
		}

		time.Sleep(time.Second)
	}
}
