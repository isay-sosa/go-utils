package utils

import (
	"fmt"
	"testing"
	"github.com/OscarSwanros/go-utils/utils"
)

type TestStruct struct {
	Value string
}

func TestIncludes(t *testing.T) {
	slice := []string{"a", "b", "c", "d"}

	include, err := Includes(slice, "c")
	if !include {
		t.Error("C element should be in collection, but it wasn't")
	}

	if err != nil {
		t.Errorf("Error should be nil, but was %s", err.Error())
	}

	include, err = Includes(slice, "z")
	if include {
		t.Error("Z element should not be in collection, but it was")
	}

	if err == nil {
		t.Error("Error should be not found, but it was nil")
	}

	include, err = Includes("Not Collection", nil)
	if include {
		t.Error("nil should no tbe in collection, but it was")
	}

	if err == nil {
		t.Error("Error should be collection values is not a slice, but it was nil")
	}
}

func TestMap(t *testing.T) {
	collection := []*TestStruct{
		&TestStruct{"Value A"},
		&TestStruct{"Value B"},
		&TestStruct{"Value C"},
	}

	newCollection, err := Map(collection, func(obj interface{}) interface{} {
		return obj.(*TestStruct).Value
	})

	if len(newCollection) != len(collection) {
		t.Errorf("New Collection len should be %d, but was %d", len(collection), len(newCollection))
	}

	if err != nil {
		t.Errorf("Error should be nil,, but was %s", err.Error())
	}

	for i := 0; i < len(newCollection); i++ {
		val := fmt.Sprintf("%v", newCollection[i])
		if val != collection[i].Value {
			t.Errorf("Value from item %d should be %s, but was %s", i, collection[i].Value, val)
		}
	}

	newCollection, err = Map("Not a Collection", nil)
	if size := len(newCollection); size != 0 {
		t.Errorf("New collection len should be 0, but was %d", size)
	}

	if err == nil || err.Error() != utils.NewError(ErrorNotASlice).Error() {
		t.Errorf("Error should be %v, but was %v", utils.NewError(ErrorNotASlice), err)
	}

	newCollection, err = Map(collection, nil)
	if size := len(newCollection); size != 0 {
		t.Errorf("New collection len should be 0, but was %d", size)
	}

	if err == nil || err.Error() != utils.NewError(ErrorNilMapFunc).Error() {
		t.Errorf("Error should be %v, but was %v", utils.NewError(ErrorNilMapFunc), err)
	}
}

func TestSelect(t *testing.T) {
	collection := []*TestStruct{
		&TestStruct{"Value A"},
		&TestStruct{"Value B"},
		&TestStruct{"Value C"},
		&TestStruct{"Value B"},
	}

	newCollection, err := Select(collection, func(obj interface{}) bool {
		return obj.(*TestStruct).Value == "Value B"
	})

	if size := len(newCollection); size != 2 {
		t.Errorf("New collection len should be 2, but was %d", size)
	}

	if err != nil {
		t.Errorf("Error should be nil,, but was %s", err.Error())
	}

	for i := 0; i < len(newCollection); i++ {
		testStruct := newCollection[i].(*TestStruct)
		if testStruct.Value != "Value B" {
			t.Errorf("Value from item %d should be 'Value B', but was %s", i, testStruct.Value)
		}
	}

	newCollection, err = Select("Not a Collection", nil)
	if size := len(newCollection); size != 0 {
		t.Errorf("New collection len should be 0, but was %d", size)
	}

	if err == nil || err.Error() != utils.NewError(ErrorNotASlice).Error() {
		t.Errorf("Error should be %v, but was %v", utils.NewError(ErrorNotASlice), err)
	}

	newCollection, err = Select(collection, nil)
	if size := len(newCollection); size != 0 {
		t.Errorf("New collection len should be 0, but was %d", size)
	}

	if err == nil || err.Error() != utils.NewError(ErrorNilSelectFunc).Error() {
		t.Errorf("Error should be %v, but was %v", utils.NewError(ErrorNilSelectFunc), err)
	}
}
