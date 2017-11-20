package gmap

type SelectFunc func(k string, v interface{}) bool

// Invokes SelectFunc for each k,v pair in the map, keeping elements for which the function returns true.
// Opposite of Reject().
func (m Map) Select(selectFn SelectFunc) Map {
	if selectFn == nil {
		return m
	}

	result := Map{}
	for k, v := range m {
		if selectFn(k, v) {
			result[k] = v
		}
	}
	return result
}
