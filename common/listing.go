package main

import (
	"encoding/json"
	"time"
)

type Listing struct {
	url              string
	date             time.Time
	city             string
	postCode         string
	lat              float32
	long             float32
	price            float32
	bedrooms         int
	rooms            int
	surface          int
	constructionYear int
}

func (listing Listing) toJson() ([]byte, error) {
	return json.Marshal(listing)
}
