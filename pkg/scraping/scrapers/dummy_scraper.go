package scrapers

import (
	. "housing_viz/pkg/common"
	"time"
)

var SampleListing = Listing{
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

func DummyScraperRun(outputChan chan<- *Listing, isDuplicate func(string) bool) error {
	time.Sleep(time.Second * 3)
	outputChan <- &SampleListing
	time.Sleep(time.Second * 4)
	outputChan <- nil
	return nil
}
