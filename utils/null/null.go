package null

import (
	"reflect"
)

func IsNil(value interface{}) bool {
	if value == nil {
		return true
	}
	if reflect.TypeOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil() {
		return true
	}
	return false
}
