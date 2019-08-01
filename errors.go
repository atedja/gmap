package gmap

import (
	"errors"
)

// ErrTypeMismatch is returned when gmap is not able to convert the underlying value to the type specified.
var ErrTypeMismatch = errors.New("gmap value type mismatch")

// ErrElementTypeMismatch is returned when one of the elements of the underlying value has a different type.
var ErrElementTypeMismatch = errors.New("gmap elements type mismatch")

// ErrKeyDoesNotExist is returned when the specified key does not exist.
var ErrKeyDoesNotExist = errors.New("gmap key does not exist")

// ErrNilValue is returned when the underlying value is nil.
var ErrNilValue = errors.New("gmap value is nil")
