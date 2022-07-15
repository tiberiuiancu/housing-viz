package main

import (
	"fmt"
	"time"
)

type DummyScraper struct {
	name    string
	running bool
}

func (s *DummyScraper) getName() string {
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

func (s *DummyScraper) isRunning() bool {
	return s.running
}
