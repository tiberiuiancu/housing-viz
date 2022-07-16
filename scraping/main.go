package main

func main() {
	Scheduler{
		scrapers: []Scraper{
			{
				name:     "Dummy",
				runnable: dummyScraperRun,
			},
		},
	}.start()
}
