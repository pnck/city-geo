package main

import (
	"encoding/json"
	citygeo "github.com/pnck/city-geo/data"
	"os"
	"testing"
)

func TestEmbed(t *testing.T) {
	f, e := os.Open("data/data.json")
	if e != nil {
		t.Fatal(e)
	}
	dj := json.NewDecoder(f)
	var cityGeoData []citygeo.GeoDataItem
	if e = dj.Decode(&cityGeoData); e != nil {
		t.Fatal(e)
	}
	if cityGeoData[0] != citygeo.Data[0] || cityGeoData[len(cityGeoData)-1] != citygeo.Data[len(citygeo.Data)-1] {
		t.Errorf("data not match")
	}
	t.Logf("itme[0]: %v", citygeo.Data[0])
}
