package gmap

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const testPayload = `
{
 "Name":"John",
 "Hash":{ "SubKey":"Value"},
 "Array":[99, 98, 97, 96, 95],
 "Value": 1,
 "Level": 464.21,
 "Flag": true,
 "StringArray": [
	"1", "a", "2"
 ],
 "MixedStringArray": [
	"1", "a", 2.9, 100, -3, "foobar"
 ],
 "TimeISO": "2017-07-10T12:13:47Z",
 "TimeRuby": "2017-07-10 12:13:47 UTC",
 "TimeRuby2": "2017-07-10 12:13:47 -0200",
 "TimeDefault": "2017-07-10 12:13:47 -0700 PDT",
 "StringAsInt": "100",
 "StringAsFloat": "100.012",
 "StringAsBool": "true"
}
`

func init() {
	fmt.Sprintln() // just so we can use fmt
}

func TestString(t *testing.T) {
	var gmap Map
	var err error
	var value string

	gmap = Map{}
	err = json.Unmarshal([]byte(testPayload), &gmap)
	assert.Nil(t, err)

	value, err = gmap.String("Name", "")
	assert.Nil(t, err)
	assert.EqualValues(t, value, "John")

	value, err = gmap.String("DoesNotExist", "")
	assert.Equal(t, ErrKeyDoesNotExist, err)
	assert.EqualValues(t, "", value)
}

func TestGMap(t *testing.T) {
	var gmap Map
	var err error
	var value Map

	gmap = Map{}
	err = json.Unmarshal([]byte(testPayload), &gmap)
	assert.Nil(t, err)

	value, err = gmap.Map("Hash", nil)
	assert.Nil(t, err)
	assert.EqualValues(t, value["SubKey"], "Value")

	value, err = gmap.Map("DoesNotExist", nil)
	assert.NotNil(t, err)
	assert.Nil(t, value)
}

func TestGMapFromObjects(t *testing.T) {
	var gmap Map
	var err error
	var value Map

	gmap = Map{
		"someKey": Map{"value": 0},
		"mapInterface": map[interface{}]interface{}{
			"value": 1,
		},
	}

	value, err = gmap.Map("someKey", nil)
	assert.Nil(t, err)
	assert.EqualValues(t, value["value"], 0)

	value, err = gmap.Map("mapInterface", nil)
	assert.Nil(t, err)
	assert.EqualValues(t, value["value"], 1)
}

func TestArray(t *testing.T) {
	var gmap Map
	var err error
	var value []interface{}

	gmap = Map{}
	err = json.Unmarshal([]byte(testPayload), &gmap)
	assert.Nil(t, err)

	value, err = gmap.Array("Array", make([]interface{}, 0))
	assert.Nil(t, err)
	assert.EqualValues(t, 99, value[0])
	assert.EqualValues(t, 98, value[1])
	assert.EqualValues(t, 97, value[2])
	assert.EqualValues(t, 96, value[3])
	assert.EqualValues(t, 95, value[4])

	value, err = gmap.Array("Does Not Exist", make([]interface{}, 0))
	assert.NotNil(t, err)
	assert.EqualValues(t, 0, len(value))
}

func TestInt(t *testing.T) {
	var gmap Map
	var err error
	var value int

	gmap = Map{}
	err = json.Unmarshal([]byte(testPayload), &gmap)
	assert.Nil(t, err)

	value, _ = gmap.Int("Value", 0)
	assert.EqualValues(t, 1, value)

	value, _ = gmap.Int("StringAsInt", 0)
	assert.EqualValues(t, 100, value)

	value, _ = gmap.Int("DoesNotExist", 9)
	assert.EqualValues(t, 9, value)
}

func TestFloat(t *testing.T) {
	var gmap Map
	var err error
	var value float64

	gmap = Map{}
	err = json.Unmarshal([]byte(testPayload), &gmap)
	assert.Nil(t, err)

	value, _ = gmap.Float("Level", 0.0)
	assert.EqualValues(t, 464.21, value)

	value, _ = gmap.Float("StringAsFloat", 0.0)
	assert.EqualValues(t, 100.012, value)

	value, _ = gmap.Float("DoesNotExist", 10.0)
	assert.EqualValues(t, 10.0, value)
}

func TestBoolean(t *testing.T) {
	var gmap Map
	var err error
	var value bool

	gmap = Map{}
	err = json.Unmarshal([]byte(testPayload), &gmap)
	assert.Nil(t, err)

	value, _ = gmap.Boolean("Flag", false)
	assert.EqualValues(t, true, value)

	value, _ = gmap.Boolean("StringAsBool", false)
	assert.EqualValues(t, true, value)

	value, _ = gmap.Boolean("DoesNotExist", false)
	assert.EqualValues(t, false, value)
}

func TestStringArray(t *testing.T) {
	var gmap Map
	var err error
	var value []string

	gmap = Map{}
	err = json.Unmarshal([]byte(testPayload), &gmap)
	assert.Nil(t, err)

	value, err = gmap.StringArray("StringArray", []string{})
	assert.Nil(t, err)
	assert.Equal(t, []string{"1", "a", "2"}, value)

	value, err = gmap.StringArray("MixedStringArray", []string{})
	assert.Nil(t, err)
	assert.Equal(t, []string{"1", "a", "2.9", "100", "-3", "foobar"}, value)
}

func TestTime(t *testing.T) {
	var gmap Map
	var err error
	var value time.Time
	var zone string

	gmap = Map{}
	err = json.Unmarshal([]byte(testPayload), &gmap)
	assert.Nil(t, err)

	def := time.Now()
	value, _ = gmap.Time("TimeISO", def)
	assert.Equal(t, time.July, value.Month())
	assert.Equal(t, 10, value.Day())
	assert.Equal(t, 2017, value.Year())
	assert.Equal(t, 12, value.Hour())
	assert.Equal(t, 13, value.Minute())
	assert.Equal(t, 47, value.Second())
	zone, _ = value.Zone()
	assert.Equal(t, "UTC", zone)

	value, _ = gmap.Time("TimeRuby", def)
	assert.Equal(t, time.July, value.Month())
	assert.Equal(t, 10, value.Day())
	assert.Equal(t, 2017, value.Year())
	assert.Equal(t, 12, value.Hour())
	assert.Equal(t, 13, value.Minute())
	assert.Equal(t, 47, value.Second())
	zone, _ = value.Zone()
	assert.Equal(t, "UTC", zone)

	value, _ = gmap.Time("TimeRuby2", def)
	assert.Equal(t, time.July, value.Month())
	assert.Equal(t, 10, value.Day())
	assert.Equal(t, 2017, value.Year())
	assert.Equal(t, 12, value.Hour())
	assert.Equal(t, 13, value.Minute())
	assert.Equal(t, 47, value.Second())

	value, _ = gmap.Time("TimeDefault", def)
	assert.Equal(t, time.July, value.Month())
	assert.Equal(t, 10, value.Day())
	assert.Equal(t, 2017, value.Year())
	assert.Equal(t, 12, value.Hour())
	assert.Equal(t, 13, value.Minute())
	assert.Equal(t, 47, value.Second())
	zone, _ = value.Zone()
	assert.Equal(t, "PDT", zone)
}

func TestTimeUTC(t *testing.T) {
	var gmap Map
	var err error
	var value time.Time
	var zone string

	gmap = Map{}
	err = json.Unmarshal([]byte(testPayload), &gmap)
	assert.Nil(t, err)

	def := time.Now()
	value, _ = gmap.TimeUTC("TimeISO", def)
	assert.Equal(t, time.July, value.Month())
	assert.Equal(t, 10, value.Day())
	assert.Equal(t, 2017, value.Year())
	assert.Equal(t, 12, value.Hour())
	assert.Equal(t, 13, value.Minute())
	assert.Equal(t, 47, value.Second())
	zone, _ = value.Zone()
	assert.Equal(t, "UTC", zone)

	value, _ = gmap.TimeUTC("TimeRuby", def)
	assert.Equal(t, time.July, value.Month())
	assert.Equal(t, 10, value.Day())
	assert.Equal(t, 2017, value.Year())
	assert.Equal(t, 12, value.Hour())
	assert.Equal(t, 13, value.Minute())
	assert.Equal(t, 47, value.Second())
	zone, _ = value.Zone()
	assert.Equal(t, "UTC", zone)

	value, _ = gmap.TimeUTC("TimeRuby2", def)
	assert.Equal(t, time.July, value.Month())
	assert.Equal(t, 10, value.Day())
	assert.Equal(t, 2017, value.Year())
	assert.Equal(t, 14, value.Hour()) // 14 not 12. move backward 2 hours
	assert.Equal(t, 13, value.Minute())
	assert.Equal(t, 47, value.Second())
	zone, _ = value.Zone()
	assert.Equal(t, "UTC", zone)

	value, _ = gmap.TimeUTC("TimeDefault", def)
	assert.Equal(t, time.July, value.Month())
	assert.Equal(t, 10, value.Day())
	assert.Equal(t, 2017, value.Year())
	assert.Equal(t, 19, value.Hour()) // 19 not 12. move backward 7 hours
	assert.Equal(t, 13, value.Minute())
	assert.Equal(t, 47, value.Second())
	zone, _ = value.Zone()
	assert.Equal(t, "UTC", zone)
}

func TestSlice(t *testing.T) {
	var gmap Map

	gmap = Map{}
	gmap["cake"] = "is a lie"
	gmap["beer"] = "free"
	gmap["count"] = 10
	mp := gmap.Slice("cake", "count")
	assert.Equal(t, "is a lie", mp["cake"])
	assert.Equal(t, nil, mp["beer"])
	assert.Equal(t, 10, mp["count"])
}

func TestExcept(t *testing.T) {
	var gmap Map

	gmap = Map{}
	gmap["cake"] = "is a lie"
	gmap["beer"] = "free"
	gmap["count"] = 10
	mp := gmap.Except("cake", "count")
	assert.Equal(t, nil, mp["cake"])
	assert.Equal(t, "free", mp["beer"])
	assert.Equal(t, nil, mp["count"])
}

func TestFromUrlValues(t *testing.T) {
	var gmap Map

	uv := url.Values{}
	uv["foo"] = []string{"bar"}
	uv["hello"] = []string{"bar", "chomp", "bit"}
	uv["nested[map]"] = []string{"what"}
	uv["nested[is]"] = []string{"it"}
	uv["nested[1]"] = []string{"this is one", "two"}
	uv["nested[even][deeper]"] = []string{"easy there"}

	gmap = Map{}
	gmap.FromUrlValues(uv)
	assert.Equal(t, "bar", gmap["foo"])
	assert.Equal(t, []string{"bar", "chomp", "bit"}, gmap["hello"])

	nestedMap := gmap["nested"].(Map)
	assert.Equal(t, "what", nestedMap["map"])
	assert.Equal(t, "it", nestedMap["is"])
	assert.Equal(t, []string{"this is one", "two"}, nestedMap["1"])

	nestedMap = nestedMap["even"].(Map)
	assert.Equal(t, "easy there", nestedMap["deeper"])
}

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

	v = int64(1000)
	i, err = interfaceToInt(v, 0)
	assert.Nil(t, err)
	assert.Equal(t, 1000, i)

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
