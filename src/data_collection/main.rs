mod scrapers;
use scrapers::dummy;
use scrapers::pararius;
mod scheduler;
use scheduler::*;
use priority_queue::PriorityQueue;

fn main() {
    let mut sch = Scheduler {
        scrapers: {
            // make scraper list
        }
    };

    sch.scrapers.push();
}
