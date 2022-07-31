package scrapers

import (
	"fmt"
	"github.com/gocolly/colly"
	. "housing_viz/common"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func getListingPrice(e *colly.HTMLElement) (int, error) {
	price, err := e.DOM.Find(".listing-detail-summary__price").First().Html()
	if err != nil {
		return 0, err
	}
	price = regexp.MustCompile("\\d*,?\\d+").FindString(price)
	price = strings.ReplaceAll(price, ",", "")
	priceInt, err := strconv.Atoi(price)
	if err != nil {
		return 0, err
	}

	return priceInt, nil
}

func getAddress(e *colly.HTMLElement) (string, string, string, string, error) {
	title, err := e.DOM.Find(".listing-detail-summary__title").First().Html()
	if err != nil {
		return "", "", "", "", err
	}

	// get information from title
	titleSplit := strings.Split(title, " ")
	listingType := titleSplit[2]
	city := titleSplit[len(titleSplit)-1]
	street := strings.Join(titleSplit[3:len(titleSplit)-2], " ")

	// get postcode from subtitle
	subtitle, err := e.DOM.Find(".listing-detail-summary__location").First().Html()
	if err != nil {
		return "", "", "", "", err
	}
	subtitle = strings.ReplaceAll(subtitle, " ", "")
	postCode := subtitle[:6]

	return listingType, city, street, postCode, nil
}

func parariusListingFromHtml(e *colly.HTMLElement) (Listing, error) {

	// get price
	price, err := getListingPrice(e)
	if err != nil {
		return sampleListing, err
	}

	// info
	bedrooms := 2
	rooms := 3
	surface := 100
	constructionYear := 1992

	// address info
	listingType, city, street, postCode, err := getAddress(e)
	if err != nil {
		return sampleListing, err
	}

	// get latitude and longitude from address
	lat, lng, err := ResolveAddressToCoordinates(postCode)
	if err != nil {
		return sampleListing, err
	}

	return Listing{
		ScraperName:      "Pararius",
		Url:              e.Request.URL.String(),
		Date:             time.Now(),
		Country:          "Netherlands",
		City:             city,
		Street:           street,
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
	err := c.Visit("https://www.pararius.com/apartment-for-rent/rotterdam/376bfeef/laan-op-zuid")
	if err != nil {
		fmt.Println(err)
	}
}
