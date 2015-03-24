package utils

import (
	"fmt"
	"reflect"
	"testing"
)

type TestStruct struct {
	Value string
}

func TestCombination(t *testing.T) {
	c, _ := Combination([]int{1, 2, 3}, []int{4, 5})
	expectCombination := []interface{}{
		[]interface{}{1, 4},
		[]interface{}{1, 5},
		[]interface{}{2, 4},
		[]interface{}{2, 5},
		[]interface{}{3, 4},
		[]interface{}{3, 5},
	}

	if !reflect.DeepEqual(c, expectCombination) {
		t.Errorf("%v is not equal to %v", c, expectCombination)
	}

	c, _ = Combination([]int{1, 2}, []int{1, 2})
	expectCombination = []interface{}{
		[]interface{}{1, 1},
		[]interface{}{1, 2},
		[]interface{}{2, 1},
		[]interface{}{2, 2},
	}

	if !reflect.DeepEqual(c, expectCombination) {
		t.Errorf("%v is not equal to %v", c, expectCombination)
	}

	c, _ = Combination([]int{1, 2}, []int{3, 4}, []int{5, 6})
	expectCombination = []interface{}{
		[]interface{}{1, 3, 5},
		[]interface{}{1, 3, 6},
		[]interface{}{1, 4, 5},
		[]interface{}{1, 4, 6},
		[]interface{}{2, 3, 5},
		[]interface{}{2, 3, 6},
		[]interface{}{2, 4, 5},
		[]interface{}{2, 4, 6},
	}

	if !reflect.DeepEqual(c, expectCombination) {
		t.Errorf("%v is not equal to %v", c, expectCombination)
	}

	c, _ = Combination([]int{1, 2})
	expectCombination = []interface{}{
		[]interface{}{1},
		[]interface{}{2},
	}

	if !reflect.DeepEqual(c, expectCombination) {
		t.Errorf("%v is not equal to %v", c, expectCombination)
	}

	c, _ = Combination([]int{})
	expectCombination = []interface{}{}

	if !reflect.DeepEqual(c, expectCombination) {
		t.Errorf("%v is not equal to %v", c, expectCombination)
	}
}

func TestCompact(t *testing.T) {
	slice := []interface{}{
		[]int{1, 4, 5},
		nil,
		[]int{2, 4, 6},
		[]int{2, 3, 5},
	}

	c, _ := Compact(slice)
	if size := len(c); size != 3 {
		t.Errorf("Slice len should be 3, but was %d", size)
	}

	for _, item := range c {
		if item == nil {
			t.Error("Item should not be nil, but it was")
		}
	}
}

func TestIsIncluded(t *testing.T) {
	slice := []string{"a", "b", "c", "d"}

	include, err := IsIncluded(slice, "c")
	if !include {
		t.Error("C element should be in collection, but it wasn't")
	}

	if err != nil {
		t.Errorf("Error should be nil, but was %s", err.Error())
	}

	include, err = IsIncluded(slice, "z")
	if include {
		t.Error("Z element should not be in collection, but it was")
	}

	if err == nil {
		t.Error("Error should be not found, but it was nil")
	}

	include, err = IsIncluded("Not Collection", nil)
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

	if err == nil || err != NotSliceErr {
		t.Errorf("Error should be %v, but was %v", NotSliceErr, err)
	}

	newCollection, err = Map(collection, nil)
	if size := len(newCollection); size != 0 {
		t.Errorf("New collection len should be 0, but was %d", size)
	}

	if err == nil || err != NilMapFuncErr {
		t.Errorf("Error should be %v, but was %v", NilMapFuncErr, err)
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

	if err == nil || err != NotSliceErr {
		t.Errorf("Error should be %v, but was %v", NotSliceErr, err)
	}

	newCollection, err = Select(collection, nil)
	if size := len(newCollection); size != 0 {
		t.Errorf("New collection len should be 0, but was %d", size)
	}

	if err == nil || err != NilSelectFuncErr {
		t.Errorf("Error should be %v, but was %v", NilSelectFuncErr, err)
	}
}
