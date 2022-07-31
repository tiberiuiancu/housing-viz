package main

import (
	"time"
)

var sampleListing = Listing{
	scraperName:      "Dummy",
	url:              "example.com/listing123",
	date:             time.Now(),
	city:             "Amsterdam",
	street:           "Street name",
	streetNumber:     "43h",
	postCode:         "1064ab",
	lat:              1.23,
	long:             1.23,
	price:            1000,
	bedrooms:         2,
	rooms:            3,
	surface:          100,
	constructionYear: 1992,
	listingType:      "apartment",
}

func dummyScraperRun(lastScraped *Listing, outputChan chan<- *Listing) {
	if lastScraped == nil {
		// this scraper was never run before; scrape everything
		time.Sleep(time.Second) // implement in the future
	}

	time.Sleep(time.Second * 3)
	outputChan <- &sampleListing
	time.Sleep(time.Second * 4)
	outputChan <- nil
}
