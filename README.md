# gmap

[![GoDoc](https://godoc.org/github.com/atedja/gmap?status.svg)](https://godoc.org/github.com/atedja/gmap) [![Build Status](https://travis-ci.org/atedja/gmap.svg?branch=master)](https://travis-ci.org/atedja/gmap)

The missing functions for `map[string]interface{}`

This package has various utility functions and wraps all the ugly details of dealing with `interface{}` type.
It performs automatic conversion of convertible types (e.g. "100" to 100), and has the ability to parse `url.Values` so it is easier to read HTTP form data.
