package main

import "fmt"

type Scheduler struct {
	scrapers []Scraper
}

func (s Scheduler) start() {
	for _, scraper := range s.scrapers {
		fmt.Println("Running ", scraper.name)
		go scraper.run()
	}
}
