package main

import (
	"fmt"
	"time"
)

type runnable interface {
	run()
	is_running() bool
	// get_cooldown()
	get_name() string
}

type scheduler struct {
	scrapers []runnable
}

func (s scheduler) start() {
	for {
		for _, scraper := range s.scrapers {
			if scraper.is_running() {
				fmt.Println("Already running")
			} else {
				fmt.Println("Running ", scraper.get_name())
				go scraper.run()
			}
		}

		time.Sleep(time.Second / 2)
	}
}
