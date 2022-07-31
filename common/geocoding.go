package common

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type mapsApiResponse struct {
	Results []struct {
		AddressComponents []struct {
			LongName  string   `json:"long_name"`
			ShortName string   `json:"short_name"`
			Types     []string `json:"types"`
		} `json:"address_components"`
		FormattedAddress string `json:"formatted_address"`
		Geometry         struct {
			Bounds struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"bounds"`
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			LocationType string `json:"location_type"`
			Viewport     struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"viewport"`
		} `json:"geometry"`
		PlaceId string   `json:"place_id"`
		Types   []string `json:"types"`
	} `json:"results"`
	Status string `json:"status"`
}

func ResolveAddressToCoordinates(address string) (float64, float64, error) {
	// make request url
	apiKey := os.Getenv("MAPS_API_KEY")
	address = strings.ReplaceAll(address, " ", "+")
	request := "https://maps.googleapis.com/maps/api/geocode/json?address=" + address + "&key=" + apiKey

	// request
	resp, err := http.Get(request)
	if err != nil {
		return 0, 0, err
	}

	// get message body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}
	strBody := string(body)

	// json decode
	var decodedResponse mapsApiResponse
	err = json.Unmarshal([]byte(strBody), &decodedResponse)
	if err != nil {
		return 0, 0, err
	}

	// try to access data
	if decodedResponse.Status != "OK" {
		return 0, 0, errors.New("Google Maps API response code " + string(decodedResponse.Status))
	}

	return decodedResponse.Results[0].Geometry.Location.Lat,
		decodedResponse.Results[0].Geometry.Location.Lng,
		nil
}
