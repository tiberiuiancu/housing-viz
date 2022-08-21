package web

import (
	"housing_viz/pkg/common"
	"reflect"
	"testing"
)

func TestListingsToMapPointConversion(t *testing.T) {
	listings := []common.Listing{
		{
			Lat:             0.12,
			Lng:             0.12,
			AddressGroup:    "1",
			NormalizedPrice: 300,
		},
		{
			Lat:             0.13,
			Lng:             0.13,
			AddressGroup:    "1",
			NormalizedPrice: 100,
		},
		{
			Lat:             0.14,
			Lng:             0.14,
			AddressGroup:    "1",
			NormalizedPrice: 200,
		},
		{
			Lat:             1.1,
			Lng:             1.1,
			AddressGroup:    "2",
			NormalizedPrice: 300,
		},
		{
			Lat:             1.2,
			Lng:             1.2,
			AddressGroup:    "2",
			NormalizedPrice: 100,
		},
	}

	expectedMapPoints := []mapPoint{
		{
			lat:    0.12,
			lng:    0.12,
			weight: 100,
		},
		{
			lat:    0.13,
			lng:    0.13,
			weight: float64(100) / float64(3),
		},
		{
			lat:    0.14,
			lng:    0.14,
			weight: float64(200) / float64(3),
		},
		{
			lat:    1.1,
			lng:    1.1,
			weight: 150,
		},
		{
			lat:    1.2,
			lng:    1.2,
			weight: 50,
		},
	}

	mapPoints := convertListingsToMapPoint(listings)
	if !reflect.DeepEqual(mapPoints, expectedMapPoints) {
		t.Error("Error: incorrect converted map points. Expected\n", expectedMapPoints, "\nInstead got:\n", mapPoint{})
	}
}
