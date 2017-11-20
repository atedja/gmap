package gmap_test

import (
	"fmt"
	"github.com/atedja/gmap"
)

func ExampleGMap_Reject() {
	var prices = gmap.GMap{}
	prices["toothpaste"] = 100
	prices["cookies"] = 80
	prices["watermelons"] = 200
	prices["vodka"] = 400

	cheap := prices.Reject(func(k string, v interface{}) bool {
		return v.(int) > 100
	})
	fmt.Println(cheap)
}
