package scrapers

import (
	. "housing_viz/common"
	"time"
)

var sampleListing = Listing{
	ScraperName: "Dummy",
	Url:         "example.com/listing123",
	Date:        time.Now(),
	Country:     "NL",
	City:        "Amsterdam",
	Street:      "Street name",
	PostCode:    "1064ab",
	Lat:         1.23,
	Lng:         1.23,
	Price:       1000,
	Bedrooms:    2,
	Rooms:       3,
	Surface:     100,
	ListingType: "apartment",
}

func DummyScraperRun(outputChan chan<- *Listing) {
	time.Sleep(time.Second * 3)
	outputChan <- &sampleListing
	time.Sleep(time.Second * 4)
	outputChan <- nil
}
