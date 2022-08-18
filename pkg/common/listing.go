package common

import (
	"time"
)

type Listing struct {
	ScraperName string    `bson:"scraper_name"`
	Url         string    `bson:"url"`
	Date        time.Time `bson:"date"`
	Country     string    `bson:"country"`
	City        string    `bson:"city"`
	Street      string    `bson:"street"`
	PostCode    string    `bson:"post_code"`
	Lat         float64   `bson:"lat"`
	Lng         float64   `bson:"lng"`
	Price       int       `bson:"price"`
	Bedrooms    int       `bson:"bedrooms"`
	Rooms       int       `bson:"rooms"`
	Surface     int       `bson:"surface"`
	ListingType string    `bson:"listing_type"`
}
