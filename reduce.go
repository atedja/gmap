package gmap

type ReduceFunc func(memo interface{}, k string, v interface{}) interface{}

// Combines all entries of map by applying an operation specified by ReduceFunc.
// For each entry, the ReduceFunc is passed a memo value from previous iteration and the key-value pair.
// The result becomes the memo value for the next iteration.
// Returns the final memo result.
func (m Map) Reduce(initial interface{}, reduceFn ReduceFunc) interface{} {
	if reduceFn == nil {
		return initial
	}

	memo := initial
	for k, v := range m {
		memo = reduceFn(memo, k, v)
	}
	return memo
}
