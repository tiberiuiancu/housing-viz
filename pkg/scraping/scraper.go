package scraping

import (
	. "housing_viz/pkg/common"
	"time"
)

type Runnable func(outputChan chan<- *Listing, isDuplicate func(string) bool) error

type Scraper struct {
	Name        string
	Exec        Runnable
	Channel     chan *Listing
	NextRunTime time.Time
	Cooldown    time.Duration
	IsRunning   bool
}

func (s *Scraper) reschedule() time.Time {
	s.IsRunning = false
	s.NextRunTime = time.Now().Local().Add(s.Cooldown)
	return s.NextRunTime
}

func (s Scraper) shouldRun() bool {
	return !s.IsRunning && s.NextRunTime.Before(time.Now())
}

func (s *Scraper) run(isDuplicate func(string) bool) {
	s.IsRunning = true
	go s.Exec(s.Channel, isDuplicate)
}
