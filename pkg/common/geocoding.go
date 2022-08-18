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
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
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

	if len(decodedResponse.Results) < 1 {
		return 0, 0, errors.New("API call returned no results")
	}

	return decodedResponse.Results[0].Geometry.Location.Lat,
		decodedResponse.Results[0].Geometry.Location.Lng,
		nil
}
