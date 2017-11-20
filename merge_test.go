package gmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	var map1 Map
	var map2 Map

	map1 = Map{}
	map1["cake"] = "is a lie"
	map1["beer"] = "free"

	map2 = Map{}
	map2["cake"] = "is tasty"
	map2["soda"] = 10

	mp := map1.Merge(map2)
	assert.Equal(t, "is tasty", mp["cake"])
	assert.Equal(t, "free", mp["beer"])
	assert.Equal(t, 10, mp["soda"])
}
