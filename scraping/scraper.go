package main

import (
	. "housing_viz/common"
	"time"
)

type Runnable func(outputChan chan<- *Listing)

type Scraper struct {
	name        string
	exec        Runnable
	channel     chan *Listing
	nextRunTime time.Time
	cooldown    time.Duration
	isRunning   bool
}

func (s *Scraper) reschedule() {
	s.isRunning = false
	s.nextRunTime = time.Now().Local().Add(s.cooldown)
}

func (s Scraper) shouldRun() bool {
	return !s.isRunning && s.nextRunTime.Before(time.Now())
}

func (s *Scraper) run() {
	s.isRunning = true
	go s.exec(s.channel)
}
