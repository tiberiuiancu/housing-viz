package common

import (
	"encoding/json"
	"time"
)

type Listing struct {
	ScraperName      string
	Url              string
	Date             time.Time
	Country          string
	City             string
	Street           string
	PostCode         string
	Lat              float64
	Lng              float64
	Price            int
	Bedrooms         int
	Rooms            int
	Surface          int
	ConstructionYear int
	ListingType      string
}

func (listing Listing) toJson() ([]byte, error) {
	return json.Marshal(listing)
}
