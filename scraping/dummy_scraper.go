package main

import "time"

func dummyScraperRun() []Listing {
	time.Sleep(time.Second)
	return []Listing{
		{
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
		},
	}
}
