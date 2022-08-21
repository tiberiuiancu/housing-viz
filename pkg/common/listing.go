package common

import (
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

type Listing struct {
	ScraperName     string    `bson:"scraper_name"`
	Url             string    `bson:"url"`
	Date            time.Time `bson:"date"`
	Country         string    `bson:"country"`
	City            string    `bson:"city"`
	Street          string    `bson:"street"`
	PostCode        string    `bson:"post_code"`
	GeocodeAddress  string    `bson:"geocode_address"`
	AddressGroup    string    `bson:"address_group"`
	Lat             float64   `bson:"lat"`
	Lng             float64   `bson:"lng"`
	Price           int       `bson:"price"`
	NormalizedPrice float64   `bson:"normalized_price"`
	Bedrooms        int       `bson:"bedrooms"`
	Rooms           int       `bson:"rooms"`
	Surface         int       `bson:"surface"`
	ListingType     string    `bson:"listing_type"`
}

func (listing Listing) IsDuplicate(db MongoConn) bool {
	return db.Exists(
		bson.D{{"url", listing.Url}},
	)
}

func (listing Listing) Sync(db MongoConn) {
	// first check not duplicate
	if listing.IsDuplicate(db) {
		return
	}

	// before sync derive additional attribute NormalizedPrice
	listing.NormalizedPrice = float64(listing.Price) / float64(listing.Surface)

	// get latitude and longitude from address
	if lat, lng, err := ResolveAddressToCoordinates(listing.GeocodeAddress); err != nil {
		log.Println("Geocoding failed:", err)
		return
	} else {
		listing.Lat = lat
		listing.Lng = lng
	}

	// try to insert into database
	_, err := db.Insert(listing)
	if err != nil {
		log.Println("Error while syncing listing", err)
	}
}
