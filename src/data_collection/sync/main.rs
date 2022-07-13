#[path = "../scrapers/dummy.rs"] mod dummy; 

fn main() {
    println!("App started");
    dummy::scrape();
}
