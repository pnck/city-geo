package data

import (
	_ "embed"
	"encoding/json"
)

type GeoDataItem struct {
	Area     string  `json:"area"`
	City     string  `json:"city"`
	Province string  `json:"province"`
	Country  string  `json:"country"`
	Lat      float64 `json:"lat,string"`
	Lng      float64 `json:"lng,string"`
}

//go:embed data.json
var jsonStr []byte

var Data []GeoDataItem

func init() {
	if err := json.Unmarshal(jsonStr, &Data); err != nil {
		panic(err)
	}
}
