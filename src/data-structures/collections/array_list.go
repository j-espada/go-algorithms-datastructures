package collections

import "fmt"

const DefaultCapacity = 20

type ArrayList struct {
	arr      []interface{}
	size     int
	capacity int
}

func CreateArrayList(capacity int) *ArrayList {
	return &ArrayList{
		arr:      make([]interface{}, capacity),
		size:     0,
		capacity: capacity,
	}
}

func CreateArrayListDefaultCapacity() *ArrayList {
	return &ArrayList{
		arr:      make([]interface{}, DefaultCapacity),
		size:     0,
		capacity: DefaultCapacity,
	}
}

// Appends the specified element to the end of this list
func (list *ArrayList) Add(item interface{}) {
	if list.size < list.capacity {
		list.arr[list.size] = item
		list.size++
	} else {
		list.capacity = list.capacity * 2
		var newArr = make([]interface{}, list.capacity)
		for i := 0; i < list.size; i++ {
			newArr[i] = list.arr[i]
		}
		newArr[list.size] = item
		list.size++
		list.arr = newArr
	}
}

// Appends all of the elements in the specified collection to the end of this list, in the order that they are returned by the specified collection's iterator
func (list *ArrayList) AddAll(collectionInterface CollectionInterface) {
	var size = collectionInterface.Size()
	for i := 0; i < size; i++ {
		list.Add(collectionInterface.Get(i))
	}
}

// Returns the index of the first occurrence of the specified element in this list, or -1 if this list does not contain the element.
func (list *ArrayList) IndexOf(item interface{}, fn equal) int {
	var arr = list.arr
	for i := 0; i < list.size; i++ {
		if fn(arr[i], item) {
			return i
		}
	}
	return -1
}

// Returns the index of the last occurrence of the specified element in this list, or -1 if this list does not contain the element.
func (list *ArrayList) LastIndexOf(item interface{}, fn equal) int {
	var arr = list.arr
	var res = -1
	for i := 0; i < list.size; i++ {
		if fn(arr[i], item) {
			res = i
		}
	}
	return res
}

//Removes all of the elements from this list
func (list *ArrayList) Clear() {
	list.size = 0
	list.capacity = DefaultCapacity
	list.arr = make([]interface{}, DefaultCapacity)
}

//Returns true if this list contains the specified element.
func (list *ArrayList) Contains(item interface{}, fn equal) bool {
	var arr = list.arr
	for i := 0; i < list.size; i++ {
		if fn(arr[i], item) {
			return true
		}
	}
	return false
}

// Returns the number of elements in this list.
func (list *ArrayList) Size() int {
	return list.size
}

// Returns the number of elements in this list.
func (list *ArrayList) Capacity() int {
	return list.capacity
}

// Removes the first occurrence of the specified element from this list, if it is present
func (list *ArrayList) Remove(item interface{}, fn equal) {
	var indexItem = list.IndexOf(item, fn)
	if indexItem != -1 {
		var newArr = make([]interface{}, list.capacity)
		var counter = 0
		for i := 0; i < list.size; i++ {
			if i != indexItem {
				newArr[counter] = list.arr[i]
				counter++
			}
		}
		list.arr = newArr
		list.size--
	}
}

// Returns the element at the specified position in this list.
func (list *ArrayList) Get(i int) interface{} {
	if i >= list.size || i < 0 {
		return nil
	}
	return list.arr[i]
}

// Returns a view of the portion of this list between the specified fromIndex, inclusive, and toIndex, exclusive.
func (list *ArrayList) SubList(fromIndex int, toIndex int) CollectionInterface {
	if fromIndex < 0 || fromIndex > toIndex || fromIndex > list.size || toIndex < fromIndex || toIndex > list.size {
		return nil
	}
	var newList = CreateArrayList(toIndex - fromIndex + 1)
	for i := fromIndex; i < toIndex; i++ {
		newList.Add(list.arr[i])
	}
	return newList
}

// Returns an array containing all of the elements in this list in proper sequence (from first to last element)
func (list *ArrayList) ToArray() []interface{} {
	var arr = make([]interface{}, list.size)
	for i := 0; i < list.size; i++ {
		arr[i] = list.arr[i]
	}
	return arr
}

// Returns the maximum element in the list or nil if list is empty
func (list *ArrayList) Max(fn compare) interface{} {
	var result = list.arr[0]
	for i := 0; i < list.size; i++ {
		var item = list.arr[i]
		if fn(item, result) > 0 {
			result = item
		}
	}
	return result
}

// Returns the minimum element in the list or nil if list is empty
func (list *ArrayList) Min(fn compare) interface{} {
	var result interface{} = nil
	for i := 0; i < list.size; i++ {
		var item = list.arr[i]
		if result == nil || fn(item, result) < 0 {
			result = item
		}
	}
	return result
}

func (list *ArrayList) String() string {
	var str = ""
	str += "Capacity: " + string(list.capacity) + "\n"
	str += "Size: " + string(list.size) + "\n"
	for i := 0; i < list.size; i++ {
		var item = list.arr[i]
		str += fmt.Sprint(item) + " "
	}
	return str
}
