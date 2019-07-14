package gmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterfaceToString(t *testing.T) {
	var v interface{}
	var s string
	var err error

	v = 10
	s, err = interfaceToString(v, "")
	assert.Nil(t, err)
	assert.Equal(t, "10", s)

	v = 10.05
	s, err = interfaceToString(v, "")
	assert.Nil(t, err)
	assert.Equal(t, "10.05", s)

	v = false
	s, err = interfaceToString(v, "")
	assert.Nil(t, err)
	assert.Equal(t, "false", s)

	v = int64(1000)
	s, err = interfaceToString(v, "")
	assert.Nil(t, err)
	assert.Equal(t, "1000", s)

	v = make([]string, 10)
	s, err = interfaceToString(v, "")
	assert.Equal(t, ErrTypeMismatch, err)
	assert.Equal(t, "", s)
}

func TestInterfaceToInt(t *testing.T) {
	var v interface{}
	var i int
	var err error

	v = "10"
	i, err = interfaceToInt(v, 0)
	assert.Nil(t, err)
	assert.Equal(t, 10, i)

	v = false
	i, err = interfaceToInt(v, 0)
	assert.Nil(t, err)
	assert.Equal(t, 0, i)

	v = true
	i, err = interfaceToInt(v, 0)
	assert.Nil(t, err)
	assert.Equal(t, 1, i)

	v = int64(-1000)
	i, err = interfaceToInt(v, 0)
	assert.Nil(t, err)
	assert.Equal(t, -1000, i)

	v = int32(-1000)
	i, err = interfaceToInt(v, 0)
	assert.Nil(t, err)
	assert.Equal(t, -1000, i)

	v = int16(-1000)
	i, err = interfaceToInt(v, 0)
	assert.Nil(t, err)
	assert.Equal(t, -1000, i)

	v = int8(-100)
	i, err = interfaceToInt(v, 0)
	assert.Nil(t, err)
	assert.Equal(t, -100, i)

	v = uint64(1000)
	i, err = interfaceToInt(v, 0)
	assert.Nil(t, err)
	assert.Equal(t, 1000, i)

	v = uint32(1000)
	i, err = interfaceToInt(v, 0)
	assert.Nil(t, err)
	assert.Equal(t, 1000, i)

	v = uint16(1000)
	i, err = interfaceToInt(v, 0)
	assert.Nil(t, err)
	assert.Equal(t, 1000, i)

	v = uint8(100)
	i, err = interfaceToInt(v, 0)
	assert.Nil(t, err)
	assert.Equal(t, 100, i)

	v = make([]string, 10)
	i, err = interfaceToInt(v, -1)
	assert.Equal(t, ErrTypeMismatch, err)
	assert.Equal(t, -1, i)
}

func TestInterfaceToFloat64(t *testing.T) {
	var v interface{}
	var f float64
	var err error

	v = "10"
	f, err = interfaceToFloat64(v, 0.0)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, f)

	v = "10.05"
	f, err = interfaceToFloat64(v, 0.0)
	assert.Nil(t, err)
	assert.Equal(t, 10.05, f)

	v = false
	f, err = interfaceToFloat64(v, 0.0)
	assert.Nil(t, err)
	assert.Equal(t, 0.0, f)

	v = true
	f, err = interfaceToFloat64(v, 0.0)
	assert.Nil(t, err)
	assert.Equal(t, 1.0, f)

	v = int64(1000)
	f, err = interfaceToFloat64(v, 0.0)
	assert.Nil(t, err)
	assert.Equal(t, 1000.0, f)

	v = make([]string, 10)
	f, err = interfaceToFloat64(v, -1.0)
	assert.Equal(t, ErrTypeMismatch, err)
	assert.Equal(t, -1.0, f)
}
