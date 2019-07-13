package gmap

import (
	"strconv"
)

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
	case int8:
		return int(v.(int8)), nil
	case int16:
		return int(v.(int16)), nil
	case int64:
		return int(v.(int64)), nil
	case uint:
		return int(v.(uint)), nil
	case uint8:
		return int(v.(uint8)), nil
	case uint16:
		return int(v.(uint16)), nil
	case uint64:
		return int(v.(uint64)), nil
	case float32:
		return int(v.(float32)), nil
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
