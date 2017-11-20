package gmap

import (
	"encoding/json"
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
 "TimeISO": "2017-07-10T12:13:47Z",
 "TimeRuby": "2017-07-10 12:13:47 UTC",
 "TimeRuby2": "2017-07-10 12:13:47 -0700",
 "TimeDefault": "2017-07-10 12:13:47 -0700 PDT"
}
`

func TestString(t *testing.T) {
	var gmap GMap
	var err error
	var value string

	gmap = GMap{}
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
	var gmap GMap
	var err error
	var value GMap

	gmap = GMap{}
	err = json.Unmarshal([]byte(testPayload), &gmap)
	assert.Nil(t, err)

	value, err = gmap.GMap("Hash", nil)
	assert.Nil(t, err)
	assert.EqualValues(t, value["SubKey"], "Value")

	value, err = gmap.GMap("DoesNotExist", nil)
	assert.NotNil(t, err)
	assert.Nil(t, value)
}

func TestArray(t *testing.T) {
	var gmap GMap
	var err error
	var value []interface{}

	gmap = GMap{}
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
	var gmap GMap
	var err error
	var value int

	gmap = GMap{}
	err = json.Unmarshal([]byte(testPayload), &gmap)
	assert.Nil(t, err)

	value, _ = gmap.Int("Value", 0)
	assert.EqualValues(t, 1, value)

	value, _ = gmap.Int("DoesNotExist", 9)
	assert.EqualValues(t, 9, value)
}

func TestFloat(t *testing.T) {
	var gmap GMap
	var err error
	var value float64

	gmap = GMap{}
	err = json.Unmarshal([]byte(testPayload), &gmap)
	assert.Nil(t, err)

	value, _ = gmap.Float("Level", 0.0)
	assert.EqualValues(t, 464.21, value)

	value, _ = gmap.Float("DoesNotExist", 10.0)
	assert.EqualValues(t, 10.0, value)
}

func TestBoolean(t *testing.T) {
	var gmap GMap
	var err error
	var value bool

	gmap = GMap{}
	err = json.Unmarshal([]byte(testPayload), &gmap)
	assert.Nil(t, err)

	value, _ = gmap.Boolean("Flag", false)
	assert.EqualValues(t, true, value)

	value, _ = gmap.Boolean("DoesNotExist", false)
	assert.EqualValues(t, false, value)
}

func TestStringArray(t *testing.T) {
	var gmap GMap
	var err error
	var value []string

	gmap = GMap{}
	err = json.Unmarshal([]byte(testPayload), &gmap)
	assert.Nil(t, err)

	value, err = gmap.StringArray("StringArray", []string{})
	assert.Nil(t, err)
	assert.Equal(t, []string{"1", "a", "2"}, value)
}

func TestTime(t *testing.T) {
	var gmap GMap
	var err error
	var value time.Time
	var zone string

	gmap = GMap{}
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
	zone, _ = value.Zone()
	assert.Equal(t, "PDT", zone)

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
	var gmap GMap
	var err error
	var value time.Time
	var zone string

	gmap = GMap{}
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
	assert.Equal(t, 19, value.Hour()) // 19 not 12. move forward 7 hours
	assert.Equal(t, 13, value.Minute())
	assert.Equal(t, 47, value.Second())
	zone, _ = value.Zone()
	assert.Equal(t, "UTC", zone)

	value, _ = gmap.TimeUTC("TimeDefault", def)
	assert.Equal(t, time.July, value.Month())
	assert.Equal(t, 10, value.Day())
	assert.Equal(t, 2017, value.Year())
	assert.Equal(t, 19, value.Hour()) // 19 not 12. move forward 7 hours
	assert.Equal(t, 13, value.Minute())
	assert.Equal(t, 47, value.Second())
	zone, _ = value.Zone()
	assert.Equal(t, "UTC", zone)
}
