use crate::scrapers::abc::*;
use std::time::{Instant, Duration};

pub struct DummyScraper {
    pub next_run: Instant,
    pub schedule_every: Duration,
}

impl Scrape for DummyScraper {
    fn scrape(&self) {
        println!("scrapin'");
    }
}

pub fn init() -> DummyScraper {
    DummyScraper {
        next_run: Instant::now(),
        schedule_every: Duration::from_secs(1),
    }
}
