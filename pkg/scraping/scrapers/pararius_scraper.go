package scrapers

import (
	"errors"
	"github.com/gocolly/colly"
	. "housing_viz/pkg/common"
	"log"
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

func getSurface(e *colly.HTMLElement) (int, error) {
	surface, err := e.DOM.Find(".listing-features__description--surface_area").First().Children().First().Html()
	if err != nil {
		return -1, err
	}

	spaceIndex := strings.Index(surface, " ")
	if spaceIndex == -1 {
		return -1, errors.New("surface area could not be retrieved")
	}

	return strconv.Atoi(surface[:spaceIndex])
}

func getNumberOfRooms(e *colly.HTMLElement) (int, error) {
	rooms, err := e.DOM.Find(".listing-features__description--number_of_rooms").First().Children().First().Html()
	if err != nil {
		return -1, err
	}

	return strconv.Atoi(rooms)
}

func getNumberOfBedrooms(e *colly.HTMLElement) (int, error) {
	rooms, err := e.DOM.Find(".listing-features__description--number_of_bedrooms").First().Children().First().Html()
	if err != nil {
		return -1, err
	}

	return strconv.Atoi(rooms)
}

func parariusListingFromHtml(e *colly.HTMLElement) (Listing, error) {

	// get price
	price, err := getListingPrice(e)
	if err != nil {
		return SampleListing, err
	}

	// surface
	surface, err := getSurface(e)
	if err != nil {
		return SampleListing, err
	}

	// rooms
	rooms, err := getNumberOfRooms(e)
	if err != nil {
		return SampleListing, err
	}

	bedrooms, err := getNumberOfBedrooms(e)
	if err != nil {
		return SampleListing, err
	}

	// address info
	listingType, city, street, postCode, err := getAddress(e)
	if err != nil {
		return SampleListing, err
	}

	// derive address group
	addressGroup := "Netherlands " + postCode[:4]

	// add geocoding address
	geocodeAddress := postCode

	return Listing{
		ScraperName:    "Pararius",
		Url:            e.Request.URL.String(),
		Date:           time.Now(),
		Country:        "Netherlands",
		City:           city,
		Street:         street,
		PostCode:       postCode,
		AddressGroup:   addressGroup,
		GeocodeAddress: geocodeAddress,
		Price:          price,
		Bedrooms:       bedrooms,
		Rooms:          rooms,
		Surface:        surface,
		ListingType:    listingType,
	}, nil
}

func ParariusScraperRun(outputChan chan<- *Listing) error {

	c := colly.NewCollector(
		colly.UserAgent("*"),
		colly.AllowedDomains("www.pararius.com"),
	)

	// if we land on a listing's page, scrape it
	c.OnHTML(".page__row--listing", func(e *colly.HTMLElement) {
		go func() {
			listing, err := parariusListingFromHtml(e)
			if err == nil {
				outputChan <- &listing
			} else {
				log.Println(err)
			}
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
		if strings.Contains(link, "-for-rent") &&
			!strings.Contains(link, "login") &&
			!strings.Contains(link, "map") &&
			!strings.Contains(link, "subscribe") {
			c.Visit(link)
		}
	})

	err := c.Visit("https://www.pararius.com")
	if err != nil {
		// write nil on the channel
		outputChan <- nil
		log.Println(err)
	}

	return err
}
