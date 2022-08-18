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

func isDuplicate(listing Listing, db MongoConn) bool {
	// basic duplication check by url
	return db.Exists(
		bson.D{{"url", listing.Url}},
	)
}

func syncListing(listing Listing, db MongoConn) {
	// only sync non-dups
	if isDuplicate(listing, db) {
		log.Println("Found duplicate", listing.Url)
		return
	}

	_, err := db.Insert(listing)
	if err != nil {
		log.Println("Error while syncing listing", err)
	}
}

func (s Scheduler) Start(db MongoConn) {
	for {
		for idx := range s.Scrapers {
			scraper := &s.Scrapers[idx]

			if scraper.shouldRun() {
				// run if necessary
				log.Println("Starting scraper", scraper.Name, "in background")
				scraper.run()
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
							go syncListing(*newListing, db)
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

				log.Println("Received", nRecv, "listings from", scraper.Name)
			}
		}

		time.Sleep(time.Second)
	}
}
