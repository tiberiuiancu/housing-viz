package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	"time"
)

func parariusListingFromHtml(e *colly.HTMLElement) Listing {
	return Listing{
		scraperName:      "Pararius",
		url:              e.Request.URL.String(),
		date:             time.Now(),
		city:             "",
		street:           "",
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
	}
	fmt.Println("new listing", e)
	return sampleListing
}

func parariusScraperRun(lastScraped *Listing, outputChan chan<- *Listing) {

	c := colly.NewCollector(
		colly.UserAgent("*"),
		colly.AllowedDomains("www.pararius.com"),
	)

	// if we land on a listing's page, scrape it
	c.OnHTML(".listing-detail-summary", func(e *colly.HTMLElement) {
		listing := parariusListingFromHtml(e)
		outputChan <- &listing
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
