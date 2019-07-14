# gmap

[![GoDoc](https://godoc.org/github.com/atedja/gmap?status.svg)](https://godoc.org/github.com/atedja/gmap) [![Build Status](https://travis-ci.org/atedja/gmap.svg?branch=master)](https://travis-ci.org/atedja/gmap)

The missing functions for `map[string]interface{}`

This package has various utility functions and wraps all the ugly details of dealing with `interface{}` type.

 Feature Summary:

* Automatic Type Conversion from various formats to `int`, `float64`, `string`, and `time.Time`.
* `string` to `time.Time` auto conversion accepts the following time formats:
  * ISO8601
  * RFC1123/RFC2822
  * [Common Log Format](https://en.wikipedia.org/wiki/Common_Log_Format)
  * Golang [`time.Time.String()`](https://golang.org/pkg/time/#Time.String) format.
  * Ruby `Time#to_s` default format.
* `Slice` and `Except` to filter out keys.
* `Select` and `Reject` to filter out key/value pairs using a custom function.
* `Reduce` to reduce your map using a custom function.
* Parse `url.Values` to make it easier to read HTTP form data. Even with nested hashes.
