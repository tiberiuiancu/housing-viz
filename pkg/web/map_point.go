package web

import "housing_viz/pkg/common"

type mapPoint struct {
	Lat    float64
	Lng    float64
	Weight float64
}

func convertListingsToMapPoint(listings []common.Listing) []mapPoint {
	addressGroupWeight := make(map[string]int)

	// first pass, calculate how many occurrences of each address group there are
	for _, listing := range listings {
		if cnt, ok := addressGroupWeight[listing.AddressGroup]; ok {
			addressGroupWeight[listing.AddressGroup] = cnt + 1
		} else {
			addressGroupWeight[listing.AddressGroup] = 1
		}
	}

	// second pass, create mapPoint list
	var mapPoints []mapPoint
	for _, listing := range listings {
		listingNormalizedWeight := listing.NormalizedPrice / float64(addressGroupWeight[listing.AddressGroup])

		mapPoints = append(mapPoints, mapPoint{
			listing.Lat,
			listing.Lng,
			listingNormalizedWeight,
		})
	}

	return mapPoints
}
