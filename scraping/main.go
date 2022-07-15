package main

func main() {
	s := scheduler{
		scrapers: []runnable{
			&DummyScraper{name: "Dummy", running: false},
		},
	}
	s.start()
}
