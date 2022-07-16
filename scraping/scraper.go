package main

type Runnable func() []Listing

type Scraper struct {
	name     string
	runnable Runnable
}

func (scraper Scraper) run() []Listing {
	return scraper.runnable()
}
