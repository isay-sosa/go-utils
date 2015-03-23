package utils

import (
	"fmt"
	"testing"
)

type TestStruct struct {
	Value string
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
