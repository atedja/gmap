package gmap

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var timeformats = []string{
	"2006-01-02T15:04:05Z",
	"2006-01-02 15:04:05 -0700",
	"2006-01-02 15:04:05 MST",
	"2006-01-02 15:04:05 -0700 MST",
}

// ErrTypeMismatch is returned when gmap is not able to convert the underlying value to the type specified.
var ErrTypeMismatch = errors.New("Key type mismatch")

// ErrElementTypeMismatch is returned when one of the elements of the underlying value has a different type.
var ErrElementTypeMismatch = errors.New("One of the elements type mismatch")

// ErrKeyDoesNotExist is returned when the specified key does not exist.
var ErrKeyDoesNotExist = errors.New("Key does not exist")

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
		return def, ErrTypeMismatch
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
		return def, ErrTypeMismatch
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
		return def, ErrTypeMismatch
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
		return def, ErrTypeMismatch
	}

	switch value.(type) {
	case uint:
		return float64(value.(uint)), nil
	case uint8:
		return float64(value.(uint8)), nil
	case uint16:
		return float64(value.(uint16)), nil
	case uint32:
		return float64(value.(uint32)), nil
	case uint64:
		return float64(value.(uint64)), nil
	case int:
		return float64(value.(int)), nil
	case int8:
		return float64(value.(int8)), nil
	case int16:
		return float64(value.(int16)), nil
	case int32:
		return float64(value.(int32)), nil
	case int64:
		return float64(value.(int64)), nil
	case float32:
		return float64(value.(float32)), nil
	case float64:
		return value.(float64), nil
	case string:
		return strconv.ParseFloat(value.(string), 64)
	default:
		return def, ErrTypeMismatch
	}
}

// Retrieves a string.
// Returns the default value and an error if key does not exist or nil.
func (m Map) String(key string, def string) (string, error) {
	value, ok := m[key]
	if !ok {
		return def, ErrKeyDoesNotExist
	}

	if value == nil {
		return def, ErrTypeMismatch
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
		return def, ErrTypeMismatch
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
		return def, ErrTypeMismatch
	}

	var err error
	var sa []string
	switch value.(type) {
	case []interface{}:
		val := value.([]interface{})
		sa = make([]string, len(val))
		for i, s := range val {
			sa[i], err = interfaceToString(s, "")
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

// Retrieves time.
// Can convert time value if it's a string and in the recognized format.
// Returns the default value and an error if key does not exist or nil.
func (m Map) Time(key string, def time.Time) (time.Time, error) {
	value, ok := m[key]
	if !ok {
		return def, ErrKeyDoesNotExist
	}

	if value == nil {
		return def, ErrTypeMismatch
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

// Helper function to convert an interface{} to string
func interfaceToString(v interface{}, def string) (string, error) {
	switch v.(type) {
	case string:
		return v.(string), nil
	case bool:
		return strconv.FormatBool(v.(bool)), nil
	case float64:
		return strconv.FormatFloat(v.(float64), 'f', -1, 64), nil
	case int64:
		return strconv.FormatInt(v.(int64), 10), nil
	case uint64:
		return strconv.FormatUint(v.(uint64), 10), nil
	case int:
		return strconv.Itoa(v.(int)), nil
	default:
		return def, ErrTypeMismatch
	}
}

// Helper function to convert an interface{} to int
func interfaceToInt(v interface{}, def int) (int, error) {
	switch v.(type) {
	case int:
		return v.(int), nil
	case int64:
		return int(v.(int64)), nil
	case float64:
		return int(v.(float64)), nil
	case string:
		return strconv.Atoi(v.(string))
	case bool:
		i := 0
		if v.(bool) {
			i = 1
		}
		return i, nil
	default:
		return def, ErrTypeMismatch
	}
}

// Helper function to convert an interface{} to float64
func interfaceToFloat64(v interface{}, def float64) (float64, error) {
	switch v.(type) {
	case uint:
		return float64(v.(uint)), nil
	case uint8:
		return float64(v.(uint8)), nil
	case uint16:
		return float64(v.(uint16)), nil
	case uint32:
		return float64(v.(uint32)), nil
	case uint64:
		return float64(v.(uint64)), nil
	case int:
		return float64(v.(int)), nil
	case int8:
		return float64(v.(int8)), nil
	case int16:
		return float64(v.(int16)), nil
	case int32:
		return float64(v.(int32)), nil
	case int64:
		return float64(v.(int64)), nil
	case float32:
		return float64(v.(float32)), nil
	case float64:
		return v.(float64), nil
	case bool:
		f := 0.0
		if v.(bool) {
			f = 1.0
		}
		return f, nil
	case string:
		return strconv.ParseFloat(v.(string), 64)
	default:
		return def, ErrTypeMismatch
	}
}
