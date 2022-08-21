package scraping

import (
	"go.mongodb.org/mongo-driver/bson"
	. "housing_viz/pkg/common"
	"log"
	"time"
)

type Scheduler struct {
	Scrapers []Scraper
}

func (s Scheduler) Start(db MongoConn) {
	for {
		for idx := range s.Scrapers {
			scraper := &s.Scrapers[idx]

			if scraper.shouldRun() {
				// run scraper
				log.Println("Starting scraper", scraper.Name, "in background")
				scraper.run(func(link string) bool {
					// check if url is already in database to avoid useless requests to geocoding API
					return db.Exists(
						bson.D{{"url", link}},
					)
				})
			} else if scraper.IsRunning {
				// count number of received records
				nRecv := 0

				// the goroutine was running; check if the run completed or collect results
				for shouldReceive := true; shouldReceive; {
					select {
					// there is something to receive
					case newListing := <-scraper.Channel:
						if newListing != nil {
							// if we receive something other than nil, add it to the sync list
							go newListing.Sync(db)
							nRecv += 1
						} else {
							// if we receive nil it's a sign the goroutine finished
							// reschedule and exit for loop
							nextRun := scraper.reschedule()
							log.Println("- scraper finished. Rescheduled for", nextRun)
							shouldReceive = false
						}
					default:
						// goroutine still running; exit for loop
						shouldReceive = false
					}
				}

				log.Println("Received", nRecv, "new listings from", scraper.Name)
			}
		}

		time.Sleep(time.Second)
	}
}
