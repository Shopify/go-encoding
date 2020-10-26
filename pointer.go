package encoding

import (
	"errors"
	"reflect"
)

var (
	ErrNotAPointer = errors.New("argument must be a pointer")
)

func isPointer(data interface{}) bool {
	switch reflect.ValueOf(data).Kind() {
	case reflect.Ptr, reflect.Interface:
		return true
	default:
		return false
	}
}
