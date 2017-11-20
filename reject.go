package gmap

type RejectFunc func(k string, v interface{}) bool

// Invokes RejectFunc for each k,v pair in the map, deleting elements for which the function returns true.
// Opposite of Select().
func (g Map) Reject(rejectFn RejectFunc) Map {
	if rejectFn == nil {
		return g
	}

	result := Map{}
	for k, v := range g {
		if !rejectFn(k, v) {
			result[k] = v
		}
	}
	return result
}
