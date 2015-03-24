package utils

import (
	"reflect"
	"errors"
)

func NewError(e string) error {
	if e != "" {
		return errors.New(e)
	}

	return nil
}

func IsSlice(i interface{}) bool {
	ival := ValueOf(i)
	return ival.Kind() == reflect.Slice
}

func ValueOf(i interface{}) reflect.Value {
	return reflect.ValueOf(i)
}

func Equal(a interface{}, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}
