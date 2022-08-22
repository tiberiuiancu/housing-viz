## Housing price visualization

Website to visualize rent prices from scraped housing websites.

### Architecture 
![Architecture](docs/diagram.png)

### Sample environment variables
    MONGO_INITDB_DATABASE=test-db
    MONGO_INITDB_ROOT_USERNAME=user
    MONGO_INITDB_ROOT_PASSWORD=pass
    MONGODB_DOCKER_URI=mongodb://user:pass@database:27017
    MONGODB_URI=mongodb://user:pass@localhost:27017
    MAPS_API_KEY=<API KEY>

The API key must have the following services enabled:
* Geocoding API
* Places API
* Maps JavaScript API

### Run

To start all components, simply run

    docker-compose up

### Screenshot

![Screenshot](docs/screenshot.png)
