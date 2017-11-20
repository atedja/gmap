package gmap_test

import (
	"fmt"
	"github.com/atedja/gmap"
)

func GallonUsage(memo interface{}, k string, v interface{}) interface{} {
	m := memo.(gmap.GMap)
	m[k] = float64(v.(int)) / 40.0
	return m
}

func Example() {
	var distances = gmap.GMap{}
	distances["Las Vegas"] = 269
	distances["San Francisco"] = 382
	distances["San Diego"] = 120
	distances["Sacramento"] = 384

	var gallons = gmap.GMap{}
	distances.Reduce(gallons, GallonUsage)
	fmt.Println(gallons)
}
