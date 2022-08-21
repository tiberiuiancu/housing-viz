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
			Lat:    0.12,
			Lng:    0.12,
			Weight: 100,
		},
		{
			Lat:    0.13,
			Lng:    0.13,
			Weight: float64(100) / float64(3),
		},
		{
			Lat:    0.14,
			Lng:    0.14,
			Weight: float64(200) / float64(3),
		},
		{
			Lat:    1.1,
			Lng:    1.1,
			Weight: 150,
		},
		{
			Lat:    1.2,
			Lng:    1.2,
			Weight: 50,
		},
	}

	mapPoints := convertListingsToMapPoint(listings)
	if !reflect.DeepEqual(mapPoints, expectedMapPoints) {
		t.Error("Error: incorrect converted map points. Expected\n", expectedMapPoints, "\nInstead got:\n", mapPoint{})
	}
}
