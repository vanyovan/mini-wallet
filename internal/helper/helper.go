package helper

import (
	"reflect"
)

func IsStructEmpty(s interface{}) bool {
	return reflect.DeepEqual(s, reflect.Zero(reflect.TypeOf(s)).Interface())
}
