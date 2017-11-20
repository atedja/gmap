package gmap_test

import (
	"fmt"
	"github.com/atedja/gmap"
)

func GallonUsage(memo interface{}, k string, v interface{}) interface{} {
	m := memo.(gmap.Map)
	m[k] = float64(v.(int)) / 40.0
	return m
}

func ExampleMap_Reduce() {
	var distances = gmap.Map{}
	distances["Las Vegas"] = 269
	distances["San Francisco"] = 382
	distances["San Diego"] = 120
	distances["Sacramento"] = 384

	gallons := distances.Reduce(gmap.Map{}, GallonUsage)
	fmt.Println(gallons)
}
