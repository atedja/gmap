package gmap

type FilterFunc func(k string, v interface{}) bool

// Invokes SelectFunc for each k,v pair in the map, keeping elements for which the function returns true.
// Opposite of Reject().
func (m Map) Select(selectFn FilterFunc) Map {
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

// Invokes RejectFunc for each k,v pair in the map, deleting elements for which the function returns true.
// Opposite of Select().
func (m Map) Reject(rejectFn FilterFunc) Map {
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
