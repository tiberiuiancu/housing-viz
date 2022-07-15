package main

import (
	"fmt"
	"time"
)

type runnable interface {
	run()
	isRunning() bool
	getName() string
}

type scheduler struct {
	scrapers []runnable
}

func (s scheduler) start() {
	for {
		for _, scraper := range s.scrapers {
			if scraper.isRunning() {
				fmt.Println("Already running")
			} else {
				fmt.Println("Running ", scraper.getName())
				go scraper.run()
			}
		}

		time.Sleep(time.Second / 2)
	}
}
