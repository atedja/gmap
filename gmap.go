package gmap

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

var timeformats = []string{
	"2006-01-02T15:04:05Z",
	"Mon, 02 Jan 2006 15:04:05 MST",
	"02/Jan/2006:15:04:05 -0700",
	"2006-01-02 15:04:05 -0700",
	"2006-01-02 15:04:05 MST",
	"2006-01-02 15:04:05 -0700 MST",
}

// Map provides various utility functions for map[string]interface{}.
type Map map[string]interface{}

// Retrieves another Map (map[string]interface{}).
// Returns the default value and an error if key does not exist or nil.
func (m Map) Map(key string, def Map) (Map, error) {
	value, ok := m[key]
	if !ok {
		return def, ErrKeyDoesNotExist
	}

	if value == nil {
		return def, ErrNilValue
	}

	switch value.(type) {
	case map[string]interface{}:
		return Map(value.(map[string]interface{})), nil
	case map[interface{}]interface{}:
		mp := Map{}
		mi := value.(map[interface{}]interface{})
		for k, v := range mi {
			ks, err := interfaceToString(k, "")
			if err != nil {
				return def, ErrTypeMismatch
			}
			mp[ks] = v
		}
		return mp, nil
	case Map:
		return value.(Map), nil
	default:
		return def, ErrTypeMismatch
	}
}

// Retrieves an array of interface{}.
// Returns the default value and an error if key does not exist or nil.
func (m Map) Array(key string, def []interface{}) ([]interface{}, error) {
	value, ok := m[key]
	if !ok {
		return def, ErrKeyDoesNotExist
	}

	if value == nil {
		return def, ErrNilValue
	}

	switch value.(type) {
	case []interface{}:
		return value.([]interface{}), nil
	default:
		return def, ErrTypeMismatch
	}
}

// Retrieves an int.
// Returns the default value and an error if key does not exist or nil.
func (m Map) Int(key string, def int) (int, error) {
	value, ok := m[key]
	if !ok {
		return def, ErrKeyDoesNotExist
	}

	if value == nil {
		return def, ErrNilValue
	}

	return interfaceToInt(value, def)
}

// Retrieves a float.
// Returns the default value and an error if key does not exist or nil.
func (m Map) Float(key string, def float64) (float64, error) {
	value, ok := m[key]
	if !ok {
		return def, ErrKeyDoesNotExist
	}

	if value == nil {
		return def, ErrNilValue
	}

	return interfaceToFloat64(value, def)
}

// Retrieves a string.
// Returns the default value and an error if key does not exist or nil.
func (m Map) String(key string, def string) (string, error) {
	value, ok := m[key]
	if !ok {
		return def, ErrKeyDoesNotExist
	}

	if value == nil {
		return def, ErrNilValue
	}

	return interfaceToString(value, def)
}

// Retrieves a boolean.
// Returns the default value and an error if key does not exist or nil.
func (m Map) Boolean(key string, def bool) (bool, error) {
	value, ok := m[key]
	if !ok {
		return def, ErrKeyDoesNotExist
	}

	if value == nil {
		return def, ErrNilValue
	}

	switch value.(type) {
	case bool:
		return value.(bool), nil
	case string:
		return strconv.ParseBool(value.(string))
	default:
		return def, ErrTypeMismatch
	}
}

// Retrieves a string array.
// Returns the default value and an error if key does not exist or nil.
func (m Map) StringArray(key string, def []string) ([]string, error) {
	value, ok := m[key]
	if !ok {
		return def, ErrKeyDoesNotExist
	}

	if value == nil {
		return def, ErrNilValue
	}

	var err error
	var sa []string
	switch value.(type) {
	case []interface{}:
		val := value.([]interface{})
		sa = make([]string, len(val))
		for i, v := range val {
			sa[i], err = interfaceToString(v, "")
			if err != nil {
				return def, ErrElementTypeMismatch
			}
		}
		return sa, nil

	case []string:
		val := value.([]string)
		sa = make([]string, len(val))
		copy(sa, val)
		return sa, nil

	default:
		return def, ErrTypeMismatch
	}
}

// Retrieves an float64 array.
// Returns the default value and an error if key does not exist or nil.
func (m Map) FloatArray(key string, def []float64) ([]float64, error) {
	value, ok := m[key]
	if !ok {
		return def, ErrKeyDoesNotExist
	}

	if value == nil {
		return def, ErrNilValue
	}

	var err error
	var fa []float64
	switch value.(type) {
	case []interface{}:
		val := value.([]interface{})
		fa = make([]float64, len(val))
		for i, v := range val {
			fa[i], err = interfaceToFloat64(v, 0.0)
			if err != nil {
				return def, ErrElementTypeMismatch
			}
		}
		return fa, nil

	case []float64:
		val := value.([]float64)
		fa = make([]float64, len(val))
		copy(fa, val)
		return fa, nil

	default:
		return def, ErrTypeMismatch
	}
}

// Retrieves an int array.
// Returns the default value and an error if key does not exist or nil.
func (m Map) IntArray(key string, def []int) ([]int, error) {
	value, ok := m[key]
	if !ok {
		return def, ErrKeyDoesNotExist
	}

	if value == nil {
		return def, ErrNilValue
	}

	var err error
	var ia []int
	switch value.(type) {
	case []interface{}:
		val := value.([]interface{})
		ia = make([]int, len(val))
		for i, v := range val {
			ia[i], err = interfaceToInt(v, 0)
			if err != nil {
				return def, ErrElementTypeMismatch
			}
		}
		return ia, nil

	case []int:
		val := value.([]int)
		ia = make([]int, len(val))
		copy(ia, val)
		return ia, nil

	default:
		return def, ErrTypeMismatch
	}
}

// Retrieves time.
// Can convert time value if it's a string and in the recognized format.
// Returns the default value and an error if key does not exist or nil.
func (m Map) Time(key string, def time.Time) (time.Time, error) {
	value, ok := m[key]
	if !ok {
		return def, ErrKeyDoesNotExist
	}

	if value == nil {
		return def, ErrNilValue
	}

	switch value.(type) {
	case time.Time:
		val := value.(time.Time)
		return val, nil

	case string:
		var t time.Time
		var err error
		for _, tf := range timeformats {
			t, err = time.Parse(tf, value.(string))
			if err == nil {
				return t, nil
			}
		}
		return t, err

	default:
		return def, ErrTypeMismatch
	}
}

// Retrieves time, but also converts to UTC.
// Can convert time value if it's a string and in the recognized format.
// Returns the default value and an error if key does not exist or nil.
func (m Map) TimeUTC(key string, def time.Time) (time.Time, error) {
	t, err := m.Time(key, def)
	return t.UTC(), err
}

// Slice returns a new Map with only the given keys.
// Opposite of Except.
func (m Map) Slice(keys ...string) Map {
	mp := Map{}
	for _, k := range keys {
		if v, ok := m[k]; ok {
			mp[k] = v
		}
	}

	return mp
}

// Except returns a new Map except the given keys.
// Opposite of Slice.
func (m Map) Except(keys ...string) Map {
	mp := Map{}
	for k, v := range m {
		mp[k] = v
	}

	for _, k := range keys {
		delete(mp, k)
	}

	return mp
}

// Fills map with values from url.Values.
// Recognizes keys that are in hash form such as 'foo[bar]' and creates a nested map.
// Single-element string arrays will be unpacked to regular strings.
func (m Map) FromUrlValues(values url.Values) {
	for k, v := range values {
		subkeys := strings.FieldsFunc(k, func(c rune) bool {
			return c == '[' || c == ']'
		})

		// normal single keys
		if len(subkeys) == 1 {
			m[subkeys[0]] = v
			if len(v) == 1 {
				m[subkeys[0]] = v[0]
			}
			continue
		}

		// for multiple keys, make sure there are nested maps
		submap := m
		lastIndex := len(subkeys) - 1
		var subkey string
		for i := 0; i < lastIndex; i++ {
			subkey = subkeys[i]
			var mp Map
			if submap[subkey] == nil {
				mp = Map{}
				submap[subkey] = mp
			} else {
				// if there already exists a key but has a different type than a map
				// then we overwrite that value and replace it with a Map
				switch submap[subkey].(type) {
				case Map:
					mp = submap[subkey].(Map)
				default:
					mp = Map{}
					submap[subkey] = mp
				}
			}
			submap = mp
		}

		submap[subkeys[lastIndex]] = v
		if len(v) == 1 {
			submap[subkeys[lastIndex]] = v[0]
		}
	}
}

// Fills map given an array of keys and values.
func (m Map) FromKeysValues(keys []string, values []interface{}) {
	length := len(keys)
	if len(values) < length {
		length = len(values)
	}

	for i := 0; i < length; i++ {
		m[keys[i]] = values[i]
	}
}

// Retrieves the values of the given keys.
// If no keys are given, returns the entire map.
func (m Map) Values(keys ...string) []interface{} {
	values := make([]interface{}, 0)
	if len(keys) == 0 {
		for _, v := range m {
			values = append(values, v)
		}
	} else {
		for _, k := range keys {
			values = append(values, m[k])
		}
	}
	return values
}

// Retrieves the keys in the map.
func (m Map) Keys() []string {
	keys := make([]string, 0)
	for k, _ := range m {
		keys = append(keys, k)
	}
	return keys
}
