use crate::scrapers::abc::*;
use std::time::{Instant, Duration};

pub struct ParariusScraper {
    pub next_run: Instant,
    pub schedule_every: Duration,
}

impl Scrape for ParariusScraper {
    fn scrape(&self) {
        println!("scrapin' pararius");
    }
}

pub fn init() -> ParariusScraper {
    ParariusScraper {
        next_run: Instant::now(),
        schedule_every: Duration::from_secs(1),
    }
}
