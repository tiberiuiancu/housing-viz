package scrapers

import (
	"fmt"
	"github.com/gocolly/colly"
	. "housing_viz/common"
	"strings"
	"time"
)

func parariusListingFromHtml(e *colly.HTMLElement) (Listing, error) {
	// info
	price := 1000
	bedrooms := 2
	rooms := 3
	surface := 100
	constructionYear := 1992
	listingType := "apartment"

	// address info
	country := "NL"
	city := "Amsterdam"
	street := "street"
	streetNumber := "43h"
	postCode := "1064ab"
	lat, lng, err := ResolveAddressToCoordinates(postCode)

	if err != nil {
		return sampleListing, err
	}

	return Listing{
		ScraperName:      "Pararius",
		Url:              e.Request.URL.String(),
		Date:             time.Now(),
		Country:          country,
		City:             city,
		Street:           street,
		StreetNumber:     streetNumber,
		PostCode:         postCode,
		Lat:              lat,
		Lng:              lng,
		Price:            price,
		Bedrooms:         bedrooms,
		Rooms:            rooms,
		Surface:          surface,
		ConstructionYear: constructionYear,
		ListingType:      listingType,
	}, nil
}

func ParariusScraperRun(outputChan chan<- *Listing) {

	c := colly.NewCollector(
		colly.UserAgent("*"),
		colly.AllowedDomains("www.pararius.com"),
	)

	// if we land on a listing's page, scrape it
	c.OnHTML(".listing-detail-summary", func(e *colly.HTMLElement) {
		go func() {
			listing, err := parariusListingFromHtml(e)
			if err == nil {
				outputChan <- &listing
			}
			// todo: log error
		}()
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// get links on the page
		link := e.Attr("href")

		// prepend website domain name if necessary
		if link[0] == '/' {
			link = "https://www.pararius.com" + link
		}

		// only look for listings or list
		if (strings.Contains(link, "apartment") || strings.Contains(link, "house")) &&
			!strings.Contains(link, "login") &&
			!strings.Contains(link, "map") &&
			!strings.Contains(link, "subscribe") {
			c.Visit(link)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL: ", r.Request.URL, " failed with response: ", r, "\nError: ", err)
	})

	//err := c.Visit("https://www.pararius.com")
	err := c.Visit("https://www.pararius.com/apartment-for-rent/rotterdam/64a01901/nico-koomanskade")
	if err != nil {
		fmt.Println(err)
	}
}
