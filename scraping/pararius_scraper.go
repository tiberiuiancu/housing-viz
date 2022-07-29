package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

func listingFromHtml(e *colly.HTMLElement) Listing {
	return sampleListing
}

func parariusScraperRun(lastScraped *Listing, outputChan chan<- *Listing) {

	c := colly.NewCollector(
		colly.UserAgent("*"),
		colly.AllowedDomains("www.pararius.com"),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// get links on the page
		link := e.Attr("href")

		// append website domain name if necessary
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

	// if we land on a listing's page, scrape it
	c.OnHTML(".listing-detail-summary", func(e *colly.HTMLElement) {
		listing := listingFromHtml(e)
		outputChan <- &listing
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL: ", r.Request.URL, " failed with response: ", r, "\nError: ", err)
	})

	err := c.Visit("https://www.pararius.com")
	if err != nil {
		fmt.Println(err)
	}
}
