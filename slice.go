package utils

import (
	"errors"
	"reflect"
)

var (
	NotSliceErr   = errors.New("collection value is not a Slice.")
	NilMapFuncErr = errors.New("map function is nil.")
)

type MapFunc func(obj interface{}) interface{}

// Map calls the specified mapFunc once for each element in the collection.
// collection -> is the slice containing all the elements.
// mapFunc -> is the functions to be called by Map. It receives each element of the collection and returns the value for the new slice.
// It returns a new slice containing the values returned by the mapFunc.
// If collection is not a slice, then NotSliceErr is returned.
// If mapFunc is nil, then NilMapFuncEerr is returned.
func Map(collection interface{}, mapFunc MapFunc) ([]interface{}, error) {
	collectionValue := reflect.ValueOf(collection)
	if collectionValue.Kind() != reflect.Slice {
		return make([]interface{}, 0), NotSliceErr
	}

	if mapFunc == nil {
		return make([]interface{}, 0), NilMapFuncErr
	}

	newColl := make([]interface{}, collectionValue.Len())

	for i := 0; i < collectionValue.Len(); i++ {
		newColl[i] = mapFunc(collectionValue.Index(i).Interface())
	}

	return newColl, nil
}
