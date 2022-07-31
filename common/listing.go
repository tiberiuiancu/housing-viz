package common

import (
	"encoding/json"
	"time"
)

type Listing struct {
	ScraperName      string
	Url              string
	Date             time.Time
	City             string
	Street           string
	StreetNumber     string
	PostCode         string
	Lat              float32
	Long             float32
	Price            float32
	Bedrooms         int
	Rooms            int
	Surface          int
	ConstructionYear int
	ListingType      string
}

func (listing Listing) toJson() ([]byte, error) {
	return json.Marshal(listing)
}
