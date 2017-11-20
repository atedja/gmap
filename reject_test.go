package gmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Pricey(k string, v interface{}) bool {
	return v.(int) > 100
}

func TestReject(t *testing.T) {
	var prices = Map{}
	prices["toothpaste"] = 100
	prices["cookies"] = 80
	prices["watermelons"] = 200
	prices["vodka"] = 400
	result := prices.Reject(Pricey)
	assert.Equal(t, 100, result["toothpaste"].(int))
	assert.Equal(t, 80, result["cookies"].(int))
	assert.Nil(t, result["vodka"])
	assert.Nil(t, result["watermelons"])
}
