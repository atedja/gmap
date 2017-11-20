package gmap

type RejectFunc func(k string, v interface{}) bool

// Invokes RejectFunc for each k,v pair in the map, deleting elements for which the function returns true.
// Opposite of Select().
func (m Map) Reject(rejectFn RejectFunc) Map {
	if rejectFn == nil {
		return m
	}

	result := Map{}
	for k, v := range m {
		if !rejectFn(k, v) {
			result[k] = v
		}
	}
	return result
}
