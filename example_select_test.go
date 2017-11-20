package gmap_test

import (
	"fmt"
	"github.com/atedja/gmap"
)

func Pricey(k string, v interface{}) bool {
	return v.(int) > 100
}

func ExampleGMap_Select() {
	var prices = gmap.GMap{}
	prices["toothpaste"] = 100
	prices["cookies"] = 80
	prices["watermelons"] = 200
	prices["vodka"] = 400

	result := prices.Select(Pricey)
	fmt.Println(result)
}
