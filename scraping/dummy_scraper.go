package main

import (
	"fmt"
	"time"
)

type DummyScraper struct {
	name    string
	running bool
}

func (s *DummyScraper) get_name() string {
	return s.name
}

func (s *DummyScraper) run() {
	s.running = true
	for i := 0; i < 10; i++ {
		fmt.Println("Doing some dummy stuff")
		time.Sleep(time.Second)
	}
	s.running = false
}

func (s *DummyScraper) is_running() bool {
	return s.running
}
