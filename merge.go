package gmap

// MergeFunc determines how to merge two maps together when there is a key collision.
// Values from both maps are passed to the function, with old represents the value from the Map being merged into, and new is the value from the other Map.
type MergeFunc func(k string, oldValue, newValue interface{}) interface{}

// Merges this map with another Map.
// Entries with key collisions are overwritten with the values from other Map.
// Returns a new Map.
func (m Map) Merge(other Map) Map {
	return m.MergeWithFunc(other, func(k string, oldValue, newValue interface{}) interface{} {
		return newValue
	})
}

// Merges this Map with another Map with a custom merge function.
// Returns a new Map.
func (m Map) MergeWithFunc(other Map, mergeFn MergeFunc) Map {
	mp := Map{}
	for k, v := range m {
		mp[k] = v
	}

	for k, v := range other {
		if mval, ok := mp[k]; ok {
			mp[k] = mergeFn(k, mval, v)
		} else {
			mp[k] = v
		}
	}

	return mp
}
