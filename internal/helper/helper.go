package helper

import (
	"reflect"
	"time"
)

func IsStructEmpty(s interface{}) bool {
	return reflect.DeepEqual(s, reflect.Zero(reflect.TypeOf(s)).Interface())
}

func ParseTime(timeDesired time.Time, layout string) (time.Time, error) {
	// Get the current time
	desiredTime, err := time.Parse(layout, timeDesired.String())
	return desiredTime, err
}
