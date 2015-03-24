package utils

import (
	"github.com/OscarSwanros/go-utils/utils"
)

const (
	ErrorNotASlice = "Collection value is not a Slice."
	ErrorNilMapFunc = "MapFunc param is nil."
	ErrorNilSelectFunc = "SelectFunc param is nil."
	ErrorElementNotFound = "The element was not found."
)

type MapFunc func(obj interface{}) interface{}
type SelectFunc func(obj interface{}) bool

// IsIncluded returns true if the specified element is present in the specified collection, otherwise returns false.
// collection -> is the slice containing all the elements.
// obj -> is the element to be seek.
// If collection is not a slice, then NotSliceErr is returned.
// If element is not found, then ElemNotFoundErr is returned.
func IsIncluded(collection interface{}, obj interface{}) (bool, error) {
	if !utils.IsSlice(collection) {
		return false, utils.NewError(ErrorNotASlice)
	}

	value := utils.ValueOf(collection)
	l := value.Len()

	for i := 0; l; i++ {
		if item := value.Index(i).Interface(); utils.Equal(item, obj) {
			return true, nil
		}
	}

	return false, utils.NewError(ErrorElementNotFound)
}

// Map calls the specified mapFunc once for each element in the collection.
// collection -> is the slice containing all the elements.
// mapFunc -> is the function to be called by Map. It receives each element of the collection and returns the value for the new slice.
// It returns a new slice containing the values returned by the mapFunc.
// If collection is not a slice, then NotSliceErr is returned.
// If mapFunc is nil, then NilMapFuncErr is returned.
func Map(collection interface{}, mapFunc MapFunc) ([]interface{}, error) {
	if !utils.IsSlice(collection) {
		return nil, utils.NewError(ErrorNotASlice)
	}

	if mapFunc == nil {
		return nil, utils.NewError(ErrorNilMapFunc)
	}

	value := utils.ValueOf(collection)
	l := value.Len()

	newColl := make([]interface{}, l)
	for i := 0; i < l; i++ {
		newColl[i] = mapFunc(value.Index(i).Interface())
	}

	return newColl, nil
}

// Select calls the specified selectFunc once for each element in the collection.
// collection -> is the slice containing all the elements.
// selectFunc -> is the function to be called by Select. It receives each element of the collection and returns bool value. If true is returned,
// the element will be added to the returned slice.
// It returns a new slice containing all elements of the collection for which the specified selectFunc returns true.
// If collection is not a slice, then NotSliceErr is returned.
// If mapFunc is nil, then NilSelectFuncErr is returned.
func Select(collection interface{}, selectFunc SelectFunc) ([]interface{}, error) {
	if !utils.IsSlice(collection) {
		return nil, utils.NewError(ErrorNotASlice)
	}

	if selectFunc == nil {
		return nil, utils.NewError(ErrorNilSelectFunc)
	}

	value := utils.ValueOf(collection)
	newColl := []interface{}{}
	for i := 0; i < value.Len(); i++ {
		if obj := value.Index(i).Interface(); selectFunc(obj) {
			newColl = append(newColl, obj)
		}
	}

	return newColl, nil
}
