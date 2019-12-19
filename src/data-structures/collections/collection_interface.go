package collections

type CollectionItem interface {
	// Compare if two CollectionItem are equal returns true if equal false otherwise
	Equals(item CollectionItem) bool
	// Compare two CollectionItems return 0 if equal, return 1 if bigger and return -1 if lower
	Compare(item CollectionItem) int
}

type CollectionInterface interface {
	// Appends the specified element to the end of this list
	Add(item CollectionItem)
	// Appends all of the elements in the specified collection to the end of this list, in the order that they are returned by the specified collection's iterator
	AddAll(collectionInterface CollectionInterface)
	// Returns the index of the first occurrence of the specified element in this list, or -1 if this list does not contain the element.
	IndexOf(item CollectionItem) int
	// Returns the index of the last occurrence of the specified element in this list, or -1 if this list does not contain the element.
	LastIndexOf(item CollectionItem) int
	//Removes all of the elements from this list
	Clear()
	//Returns true if this list contains the specified element.
	Contains(item CollectionItem) bool
	// Returns the number of elements in this list.
	Size() int
	// Removes the first occurrence of the specified element from this list, if it is present
	Remove(item CollectionItem)
	// Returns the element at the specified position in this list.
	Get(i int) CollectionItem
	// Returns a view of the portion of this list between the specified fromIndex, inclusive, and toIndex, exclusive.
	SubList(fromIndex int, toIndex int) CollectionInterface
	// Returns an array containing all of the elements in this list in proper sequence (from first to last element)
	ToArray() []CollectionItem
	// Returns the maximum element in the list
	Max() CollectionItem
	// Returns the minimum element in the list
	Min() CollectionItem
}