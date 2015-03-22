package arraylist

import (
	"errors"
	"fmt"
	"reflect"
)

type ArrayList struct {
	slice []interface{}
}

// Add appends the specified element to the end of this list.
// obj -> is the element to be appended to this list.
func (a *ArrayList) Add(obj interface{}) {
	a.slice = append(a.slice, obj)
}

// AddAt inserts the specified element at the specified position in this list.
// pos -> is the position at which the specified element is to be inserted (0-based).
// obj -> is the element to be appended to this list.
// If pos is more than the list size or less than 0, then index out of range
// error is returned. Nil otherwise.
func (a *ArrayList) AddAt(pos int, obj interface{}) error {
	if err := a.checkRangeForAddAt(pos); err != nil {
		return err
	}

	switch pos {
	case 0:
		a.AddFirst(obj)
		break
	case a.Size():
		a.Add(obj)
		break
	default:
		a.addAt(pos, obj)
	}

	return nil
}

// AddFirst inserts the specified element to the begining of this list.
// obj -> is the element to be appended to this list.
func (a *ArrayList) AddFirst(obj interface{}) {
	a.slice = append([]interface{}{obj}, a.slice...)
}

// AddSlice appends all of the elements in the specified slice to the end of this list.
// slice -> is the slice containing elements to be added to this list.
func (a *ArrayList) AddSlice(slice []interface{}) {
	a.slice = append(a.slice, slice...)
}

// AddSliceFirst inserts all of the elements in the specified slice to the begining of this list.
// slice -> is the slice containing elements to be added to this list.
func (a *ArrayList) AddSliceFirst(slice []interface{}) {
	a.slice = append(slice, a.slice...)
}

// AddSliceAt inserts all of the elements in the specified slice at the specified position in this list.
// pos -> is the position at which all of the elements in the specified slice are to be inserted (0-based).
// slice -> is the slice containing elements to be added to this list.
// If pos is more than the list size or less than 0, then index out of range
// error is returned. Nil otherwise.
func (a *ArrayList) AddSliceAt(pos int, slice []interface{}) error {
	if err := a.checkRangeForAddAt(pos); err != nil {
		return err
	}

	switch pos {
	case 0:
		a.AddSliceFirst(slice)
		break
	case a.Size():
		a.AddSlice(slice)
		break
	default:
		a.addSliceAt(pos, slice)
	}

	return nil
}

// Clear removes all of the elements from this list.
func (a *ArrayList) Clear() {
	a.slice = []interface{}{}
}

// Get returns the element at the specified position in this list.
// It returns the element at the specified position if exists, otherwise returns nil.
// Can return index out of range error.
func (a *ArrayList) Get(pos int) (interface{}, error) {
	if err := a.checkRange(pos); err != nil {
		return nil, indexOutOfRangeErr(pos, a.Size())
	}

	return a.slice[pos], nil
}

// IndexOf returns the index (0-based) of the first occurrence of the specified element in this list.
// It can return -1 if this list does not contain the specified element.
func (a *ArrayList) IndexOf(obj interface{}) int {
	pos := -1

	for i, o := range a.slice {
		if reflect.DeepEqual(o, obj) {
			return i
		}
	}

	return pos
}

// IsEmpty returns true if this list containes no elements.
func (a *ArrayList) IsEmpty() bool {
	return a.Size() == 0
}

// LastIndexOf returns the index (0-based) of the last occurrence of the specified element in this list.
// It can return -1 if this list does not contain the specified element.
func (a *ArrayList) LastIndexOf(obj interface{}) int {
	pos := -1

	for i := a.Size() - 1; i > -1; i-- {
		if o := a.slice[i]; reflect.DeepEqual(o, obj) {
			return i
		}
	}

	return pos
}

// Remove removes the first occurrence of the specified element from this list.
// If element not found, it returns an element not found error.
func (a *ArrayList) Remove(obj interface{}) error {
	for i, o := range a.slice {
		if reflect.DeepEqual(o, obj) {
			return a.RemoveAt(i)
		}
	}

	return elementNotFoundErr(obj)
}

// RemoveAt removes the element at the specified position (0-based) in this list.
// It can return index out of range error.
func (a *ArrayList) RemoveAt(pos int) error {
	if err := a.checkRange(pos); err != nil {
		return err
	}

	a.slice = append(a.slice[:pos], a.slice[pos+1:]...)
	return nil
}

// Size returns the number of elements in this list.
func (a *ArrayList) Size() int {
	return len(a.slice)
}

// Slice returns a slice containing all of the elements in this list.
// To avoid references, the returned slice is a copy of this list.
func (a *ArrayList) Slice() []interface{} {
	return append([]interface{}{}, a.slice...)
}

func (a *ArrayList) addAt(pos int, obj interface{}) {
	tempSlice := []interface{}{}
	tempSlice = append(tempSlice, a.slice[:pos]...)
	tempSlice = append(tempSlice, obj)
	a.slice = append(tempSlice, a.slice[pos:]...)
}

func (a *ArrayList) addSliceAt(pos int, slice []interface{}) {
	tempSlice := []interface{}{}
	tempSlice = append(tempSlice, a.slice[:pos]...)
	tempSlice = append(tempSlice, slice...)
	a.slice = append(tempSlice, a.slice[pos:]...)
}

func (a *ArrayList) checkRangeForAddAt(pos int) error {
	if pos > a.Size() || pos < 0 {
		return indexOutOfRangeErr(pos, a.Size())
	}

	return nil
}

func (a *ArrayList) checkRange(pos int) error {
	if pos > a.Size()-1 || pos < 0 {
		return indexOutOfRangeErr(pos, a.Size())
	}

	return nil
}

func elementNotFoundErr(obj interface{}) error {
	errStr := fmt.Sprintf("%v element was not found in this list.", obj)
	return errors.New(errStr)
}

func indexOutOfRangeErr(pos, listSize int) error {
	errStr := fmt.Sprintf("Index %d is out of range from a list size of %d", pos, listSize)
	return errors.New(errStr)
}
