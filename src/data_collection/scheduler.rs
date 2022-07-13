use priority_queue::PriorityQueue;
use crate::scrapers::abc::Scrape;
use std::time::Duration;

pub struct Scheduler {
    scrapers: // priority queue here
}

impl Scheduler {
    pub fn run(&self) {
        for s in self.scrapers.iter() {
            s.scrape();
        }
    }
}
