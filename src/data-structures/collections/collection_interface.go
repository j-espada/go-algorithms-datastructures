package collections

type equal func(obj1 interface{}, obj2 interface{}) bool
type compare func(obj1 interface{}, obj2 interface{}) int

type CollectionInterface interface {
	// Appends the specified element to the end of this list
	Add(item interface{})
	// Appends all of the elements in the specified collection to the end of this list, in the order that they are returned by the specified collection's iterator
	AddAll(collectionInterface CollectionInterface)
	// Returns the index of the first occurrence of the specified element in this list, or -1 if this list does not contain the element.
	IndexOf(item interface{}, fn equal) int
	// Returns the index of the last occurrence of the specified element in this list, or -1 if this list does not contain the element.
	LastIndexOf(item interface{}, fn equal) int
	//Removes all of the elements from this list
	Clear()
	//Returns true if this list contains the specified element.
	Contains(item interface{}, fn equal) bool
	// Returns the number of elements in this list.
	Size() int
	// Removes the first occurrence of the specified element from this list, if it is present
	Remove(item interface{}, fn equal)
	// Returns the element at the specified position in this list.
	Get(i int) interface{}
	// Returns a view of the portion of this list between the specified fromIndex, inclusive, and toIndex, exclusive.
	SubList(fromIndex int, toIndex int) CollectionInterface
	// Returns an array containing all of the elements in this list in proper sequence (from first to last element)
	ToArray() []interface{}
	// Returns the maximum element in the list or nil if list is empty
	Max(fn compare) interface{}
	// Returns the minimum element in the list or nil if list is empty
	Min(fn compare) interface{}
}
