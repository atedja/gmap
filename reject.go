package gmap

type RejectFunc func(k string, v interface{}) bool

// Invokes RejectFunc for each k,v pair in the map, deleting elements for which the function returns true.
// Opposite of Select().
func (g GMap) Reject(rejectFn RejectFunc) GMap {
	if rejectFn == nil {
		return g
	}

	result := GMap{}
	for k, v := range g {
		if !rejectFn(k, v) {
			result[k] = v
		}
	}
	return result
}
