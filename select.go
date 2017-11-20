package gmap

type SelectFunc func(k string, v interface{}) bool

// Invokes `SelectFunc` for each k,v pair in the map, keeping elements for which
// the function returns true. Opposite of Reject().
func (g GMap) Select(selectFn SelectFunc) GMap {
	if selectFn == nil {
		return g
	}

	result := GMap{}
	for k, v := range g {
		if selectFn(k, v) {
			result[k] = v
		}
	}
	return result
}
