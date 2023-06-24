package helper

import (
	"reflect"

	"github.com/google/uuid"
)

func IsStructEmpty(s interface{}) bool {
	return reflect.DeepEqual(s, reflect.Zero(reflect.TypeOf(s)).Interface())
}

func GenerateGuuid() string {
	uuid := uuid.New()
	return uuid.String()
}
