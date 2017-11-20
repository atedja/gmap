package gmap_test

import (
	"fmt"
	"github.com/atedja/gmap"
)

func ExampleGMap_Select() {
	var prices = gmap.Map{}
	prices["toothpaste"] = 100
	prices["cookies"] = 80
	prices["watermelons"] = 200
	prices["vodka"] = 400

	cheap := prices.Select(func(k string, v interface{}) bool {
		return v.(int) < 100
	})
	fmt.Println(cheap)
}
