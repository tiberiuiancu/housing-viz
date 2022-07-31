package main

import (
	"encoding/json"
	"time"
)

type Listing struct {
	scraperName      string
	url              string
	date             time.Time
	city             string
	street           string
	streetNumber     string
	postCode         string
	lat              float32
	long             float32
	price            float32
	bedrooms         int
	rooms            int
	surface          int
	constructionYear int
	listingType      string
}

func (listing Listing) toJson() ([]byte, error) {
	return json.Marshal(listing)
}
