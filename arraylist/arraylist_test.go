package arraylist

import (
	"fmt"
	"testing"
)

func TestAdd_Single(t *testing.T) {
	list := new(ArrayList)

	list.Add("First Element")

	if size := list.Size(); size != 1 {
		t.Errorf("ArrayList should have a size of 1, but has %d", size)
	}

	list.Add("Second Element")
	if size := list.Size(); size != 2 {
		t.Errorf("ArrayList should have a size of 2, but has %d", size)
	}
}

func TestAdd_Several(t *testing.T) {
	slice := make([]interface{}, 10)
	for i := 0; i < 10; i++ {
		slice[i] = fmt.Sprintf("Element %d", i)
	}

	list := new(ArrayList)

	if size := list.Size(); size != 0 {
		t.Errorf("ArrayList should have a size of 0, but has %d", size)
	}

	list.Add(slice...)
	if size := list.Size(); size != 10 {
		t.Errorf("ArrayList should have a size of 10, but has %d", size)
	}
}

func TestAddAt_Single(t *testing.T) {
	list := new(ArrayList)
	for i := 0; i < 10; i++ {
		list.Add(fmt.Sprintf("Element %d", i))
	}

	list.AddAt(7, "Inserted Element")
	if size := list.Size(); size != 11 {
		t.Errorf("ArrayList should have a size of 11, but has %d", size)
	}

	if obj, _ := list.Get(7); obj != "Inserted Element" {
		t.Errorf("ArrayList 7th element should be 'Inserted Element', but was '%s'", obj)
	}

	if err := list.AddAt(20, "Not Inserted"); err == nil {
		t.Error("Error should be index out of range")
	}

	list.AddAt(0, "First Element")
	if size := list.Size(); size != 12 {
		t.Errorf("ArrayList should have a size of 12, but has %d", size)
	}

	if obj, _ := list.Get(0); obj != "First Element" {
		t.Errorf("ArrayList 0 element should be 'First Element', but was '%s'", obj)
	}

	list.AddAt(list.Size(), "Last Element")
	if size := list.Size(); size != 13 {
		t.Errorf("ArrayList should have a size of 13, but has %d", size)
	}

	if obj, _ := list.Get(12); obj != "Last Element" {
		t.Errorf("ArrayList last element should be 'Last Element', but was '%s'", obj)
	}
}

func TestAddAt_Several(t *testing.T) {
	list := new(ArrayList)
	for i := 0; i < 10; i++ {
		list.Add(fmt.Sprintf("Element %d", i))
	}

	list.AddAt(7, "Element 10", "Element 11", "Element 12")
	if size := list.Size(); size != 13 {
		t.Errorf("ArrayList should have a size of 13, but has %d", size)
	}

	for i := 0; i < 3; i++ {
		expectedObj := fmt.Sprintf("Element %d", 10+i)
		if obj, _ := list.Get(7 + i); obj != expectedObj {
			t.Errorf("ArrayList %dth element should be '%s' but was '%s'", i+7, expectedObj, obj)
		}
	}

	if err := list.AddAt(20, "Not Inserted"); err == nil {
		t.Error("Error should be index out of range")
	}

	list.AddAt(0, "First Element")
	if size := list.Size(); size != 14 {
		t.Errorf("ArrayList should have a size of 14, but has %d", size)
	}

	if obj, _ := list.Get(0); obj != "First Element" {
		t.Errorf("ArrayList 0 element should be 'First Element', but was '%s'", obj)
	}

	list.AddAt(list.Size(), "Last Element")
	if size := list.Size(); size != 15 {
		t.Errorf("ArrayList should have a size of 15, but has %d", size)
	}

	if obj, _ := list.Get(14); obj != "Last Element" {
		t.Errorf("ArrayList last element should be 'Last Element', but was '%s'", obj)
	}
}

func TestAddFirst_Single(t *testing.T) {
	list := new(ArrayList)
	for i := 1; i < 10; i++ {
		list.Add(fmt.Sprintf("Element %d", i))
	}

	list.AddFirst("Element 0")
	if size := list.Size(); size != 10 {
		t.Errorf("ArrayList should have a size of 10, but has %d", size)
	}

	if obj, _ := list.Get(0); obj != "Element 0" {
		t.Errorf("ArrayList 0 element should be 'First Element', but was '%s'", obj)
	}
}

func TestAddFirst_Several(t *testing.T) {
	list := new(ArrayList)
	for i := 10; i < 20; i++ {
		list.Add(fmt.Sprintf("Element %d", i))
	}

	if size := list.Size(); size != 10 {
		t.Errorf("ArrayList should have a size of 10, but has %d", size)
	}

	slice := make([]interface{}, 10)
	for i := 0; i < 10; i++ {
		slice[i] = fmt.Sprintf("Element %d", i)
	}

	list.AddFirst(slice...)

	if size := list.Size(); size != 20 {
		t.Errorf("ArrayList should have a size of 20, but has %d", size)
	}

	for i := 0; i < 20; i++ {
		expectedObj := fmt.Sprintf("Element %d", i)
		if obj, _ := list.Get(i); obj != expectedObj {
			t.Errorf("%v was expected to be equals to %v", obj, expectedObj)
		}
	}
}

func TestClear(t *testing.T) {
	list := new(ArrayList)
	for i := 0; i < 10; i++ {
		list.Add(fmt.Sprintf("Element %d", i))
	}

	list.Clear()
	if size := list.Size(); size != 0 {
		t.Errorf("ArrayList should have a size of 0, but has %d", size)
	}

	list.Add(1, 2, 3)
	if size := list.Size(); size != 3 {
		t.Errorf("ArrayList should have a size of 3, but has %d", size)
	}
}

func TestGet(t *testing.T) {
	list := new(ArrayList)
	for i := 0; i < 10; i++ {
		list.Add(fmt.Sprintf("Element %d", i))
	}

	obj, err := list.Get(-1)
	if obj != nil {
		t.Errorf("Element at position -1 should be nil, but was %s", obj)
	}

	if err == nil {
		t.Error("Error should be index out of range")
	}

	obj, err = list.Get(0)
	if obj == nil {
		t.Error("Element at position 0 should not be nil, but it was")
	}

	if err != nil {
		t.Errorf("Error should be nil, but was %s", err.Error())
	}

	obj, err = list.Get(list.Size())
	if obj != nil {
		t.Errorf("Element should be nil, but was %s", obj)
	}

	if err == nil {
		t.Error("Error should be index out of range")
	}
}

func TestIndexOf(t *testing.T) {
	list := new(ArrayList)
	for i := 0; i < 10; i++ {
		list.Add(fmt.Sprintf("Element %d", i))
	}

	if indexOf := list.IndexOf("Element 4"); indexOf != 4 {
		t.Errorf("Index of element should be 4, but was %d", indexOf)
	}

	if indexOf := list.IndexOf("Unexisted Element"); indexOf != -1 {
		t.Errorf("Index of element should be -1, but was %d", indexOf)
	}
}

func TestLastIndexOf(t *testing.T) {
	list := new(ArrayList)
	slice := []interface{}{"Repeated Element", "Second Element", "Repeated Element", "Last Element"}
	list.Add(slice...)

	if i := list.LastIndexOf("Repeated Element"); i != 2 {
		t.Errorf("Last index of element should be 2, but was %d", i)
	}

	if i := list.LastIndexOf("Unexisted Element"); i != -1 {
		t.Errorf("Last index of element should be -1, but was %d", i)
	}
}

func TestRemove(t *testing.T) {
	list := new(ArrayList)
	for i := 0; i < 10; i++ {
		list.Add(fmt.Sprintf("Element %d", i))
	}

	if err := list.Remove("Unexisted Element"); err == nil {
		t.Error("Error should be element not found")
	}

	if size := list.Size(); size != 10 {
		t.Errorf("ArrayList should have a size of 10, but has %d", size)
	}

	if err := list.Remove("Element 5"); err != nil {
		t.Errorf("Error should be nil, but was %s", err.Error())
	}

	if size := list.Size(); size != 9 {
		t.Errorf("ArrayList should have a size of 9, but has %d", size)
	}
}

func TestRemoveAt(t *testing.T) {
	list := new(ArrayList)
	for i := 0; i < 10; i++ {
		list.Add(fmt.Sprintf("Element %d", i))
	}

	err := list.RemoveAt(-1)
	if err == nil {
		t.Error("Error should be index out of range")
	}

	if size := list.Size(); size != 10 {
		t.Errorf("ArrayList should have a size of 10, but has %d", size)
	}

	err = list.RemoveAt(4)
	if err != nil {
		t.Errorf("Error should be nil, but was %s", err.Error())
	}

	if size := list.Size(); size != 9 {
		t.Errorf("ArrayList should have a size of 9, but has %d", size)
	}

	err = list.RemoveAt(list.Size())
	if err == nil {
		t.Error("Error should be index out of range")
	}

	if size := list.Size(); size != 9 {
		t.Errorf("ArrayList should have a size of 9, but has %d", size)
	}
}

func TestSlice(t *testing.T) {
	list := new(ArrayList)
	for i := 0; i < 10; i++ {
		list.Add(fmt.Sprintf("Element %d", i))
	}

	listArray := list.Slice()
	if len(listArray) != list.Size() {
		t.Errorf("New Slice size should be %d, but was", list.Size(), len(listArray))
	}

	listArray[0] = "New Element"
	if obj, _ := list.Get(0); obj == "New Element" {
		t.Errorf("%s should be different from 'New Element'", obj)
	}
}
