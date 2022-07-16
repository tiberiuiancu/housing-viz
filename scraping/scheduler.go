package main

import (
	"fmt"
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
			fmt.Println(scraper.name)

			if scraper.shouldRun() {
				// run if necessary
				scraper.run(sampleListing)
			} else if scraper.isRunning {
				// the goroutine was running; check if the run completed or collect results
				for shouldReceive := true; shouldReceive; {
					select {
					// there is something to receive
					case newListing := <-scraper.channel:
						if newListing != nil {
							// if we receive something other than nil, add it to the sync list
							fmt.Println("received", newListing)
							syncQueue = append(syncQueue, *newListing)
						} else {
							// if we receive nil it's a sign the goroutine finished
							// reschedule and exit for loop
							fmt.Println("received nil")
							scraper.reschedule()
							shouldReceive = false
						}
					default:
						// goroutine still running; exit for loop
						fmt.Println("still running")
						shouldReceive = false
					}
				}
			} else {
				fmt.Println("in cooldown")
			}
		}

		time.Sleep(time.Second)
	}
}
