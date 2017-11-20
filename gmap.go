package gmap

import (
	"errors"
	"time"
)

var timeformats = []string{
	"2006-01-02T15:04:05Z",
	"2006-01-02 15:04:05 -0700",
	"2006-01-02 15:04:05 MST",
	"2006-01-02 15:04:05 -0700 MST",
}

var ErrTypeMismatch = errors.New("Key type mismatch")
var ErrElementTypeMismatch = errors.New("One of the elements type mismatch")
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

	switch value.(type) {
	case int:
		return value.(int), nil
	case int64:
		return int(value.(int64)), nil
	case float64:
		return int(value.(float64)), nil
	default:
		return def, ErrTypeMismatch
	}
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
	case float64:
		return value.(float64), nil
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

	switch value.(type) {
	case string:
		return value.(string), nil
	default:
		return def, ErrTypeMismatch
	}
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

	var sa []string
	switch value.(type) {
	case []interface{}:
		val := value.([]interface{})
		sa = make([]string, len(val))
		for i, s := range val {
			sa[i], ok = s.(string)
			if !ok {
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

// Slices this map to return a new Map with only the given keys.
func (m Map) Slice(keys ...string) Map {
	mp := Map{}
	for _, k := range keys {
		if v, ok := m[k]; ok {
			mp[k] = v
		}
	}

	return mp
}

// Returns a new Map except the given keys.
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
