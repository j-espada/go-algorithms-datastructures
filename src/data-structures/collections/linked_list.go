package collections

import (
	"fmt"
)

type Node struct {
	Value CollectionItem
	next  *Node
}

func createNode(item CollectionItem) *Node {

	return &Node{
		item,
		nil,
	}
}

func (node *Node) String() string {

	return fmt.Sprint("Value: ", node.Value.String())
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

// Create a linked list
func CreateLinkedList() *LinkedList {

	return &LinkedList{
		nil,
		nil,
		0,
	}
}

// Appends the specified element to the end of this list
func (list *LinkedList) Add(item CollectionItem) {
	var newTail = createNode(item)
	if list.head == nil {
		list.head = newTail
		list.tail = newTail
		list.size++
	} else {
		var tail = list.tail
		tail.next = newTail
		list.tail = newTail
		list.size++
	}
}

// Appends all of the elements in the specified collection to the end of this list, in the order that they are returned by the specified collection's iterator
func (list *LinkedList) AddAll(collectionInterface CollectionInterface) {

	var size = collectionInterface.Size()
	for i := 0; i < size; i++ {
		list.Add(collectionInterface.Get(i))
	}
}

// Returns the index of the first occurrence of the specified element in this list, or -1 if this list does not contain the element.
func (list *LinkedList) IndexOf(item CollectionItem) int {

	var i = 0
	var currentNode = list.head
	var elm CollectionItem = nil

	for currentNode != nil {

		if item.Equals(currentNode.Value) {
			elm = item
			break
		}

		currentNode = currentNode.next
		i++
	}

	if elm == nil {
		return -1
	} else {
		return i
	}
}

// Returns the index of the last occurrence of the specified element in this list, or -1 if this list does not contain the element.
func (list *LinkedList) LastIndexOf(item CollectionItem) int {

	var i, j = 0, 0
	var currentNode = list.head
	var elm CollectionItem = nil

	for currentNode != nil {

		if item.Equals(currentNode.Value) {
			elm = item
			j = i
		}

		currentNode = currentNode.next
		i++
	}

	if elm == nil {
		return -1
	} else {
		return j
	}
}

//Removes all of the elements from this list
func (list *LinkedList) Clear() {

	list.head = nil
	list.tail = nil
	list.size = 0
}

//Returns true if this list contains the specified element.
func (list *LinkedList) Contains(item CollectionItem) bool {

	var currentNode = list.head
	for currentNode != nil {
		if item.Equals(currentNode.Value) {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// Returns the number of elements in this list.
func (list *LinkedList) Size() int {

	return list.size
}

// Removes the first occurrence of the specified element from this list, if it is present
func (list *LinkedList) Remove(item CollectionItem) {

	var current = list.head
	var previous *Node = nil
	var next *Node = nil
	var found = false
	for current != nil {
		if item.Equals(current.Value) {
			found = true
			break
		}
		previous = current
		current = current.next
		// in case that we have reached the end of the list without finding the item to remove
		if current != nil {
			next = current.next
		} else {
			next = nil
		}
	}

	if found == false {
		return
	}

	if (current != nil && current != list.head) && previous != nil && next != nil {
		previous.next = next
	} else if (current != nil && current != list.tail) && previous != nil && next == nil {
		previous.next = nil
	} else if previous == nil && current == list.head && next == nil {
		list.head = next
	} else if current == list.tail && previous != nil && next == nil {
		list.tail = previous
		list.tail.next = nil
	}
	list.size--
}

// Returns the element at the specified position in this list.
func (list *LinkedList) Get(i int) CollectionItem {

	if i > list.size || i < 0 {
		return nil
	}

	var currentNode = list.head
	var k = 0
	var item CollectionItem = nil
	for currentNode != nil {
		if i == k {
			item = currentNode.Value
			break
		}
		currentNode = currentNode.next
		k++
	}

	return item
}

// Returns the element at the specified position in this list.
func (list *LinkedList) GetLinkedListNode(i int) *Node {

	if i > list.size || i < 0 {
		return nil
	}

	var currentNode = list.head
	var k = 0
	var node *Node = nil
	for currentNode != nil {
		if i == k {
			node = currentNode
			break
		}
		currentNode = currentNode.next
		k++
	}

	return node
}

// Returns a view of the portion of this list between the specified fromIndex, inclusive, and toIndex, exclusive.
func (list *LinkedList) SubList(fromIndex int, toIndex int) CollectionInterface {

	if fromIndex < 0 || fromIndex > toIndex || fromIndex > list.size || toIndex < fromIndex || toIndex > list.size {
		return nil
	}

	var subList = CreateLinkedList()
	var i = 0
	var currentNode = list.head

	for currentNode != nil {

		if i >= fromIndex && i < toIndex {
			subList.Add(currentNode.Value)
		} else if i > toIndex {
			break
		}
		i++
		currentNode = currentNode.next
	}

	return subList
}

// Returns an array containing all of the elements in this list in proper sequence (from first to last element)
func (list *LinkedList) ToArray() []CollectionItem {

	var arr = make([]CollectionItem, list.size)
	var k = 0
	var currentNode = list.head
	for currentNode.next != nil {
		arr[k] = currentNode.Value
		currentNode = currentNode.next
		k++
	}
	return arr
}

// Returns the maximum element in the list or nil if list is empty
func (list *LinkedList) Max() CollectionItem {

	if list.size == 0 {
		return nil
	} else {

		var currentNode = list.head
		var max CollectionItem = nil

		for currentNode.next != nil {

			if max == nil || max.Compare(currentNode.Value) > 0 {
				max = currentNode.Value
			}
			currentNode = currentNode.next
		}

		return max
	}
}

// Returns the minimum element in the list or nil if list is empty
func (list *LinkedList) Min() CollectionItem {

	if list.head == nil {
		return nil
	} else {
		var currentNode = list.head
		var min CollectionItem = nil

		for currentNode.next != nil {

			if min == nil || min.Compare(currentNode.Value) < 0 {
				min = currentNode.Value
			}
			currentNode = currentNode.next
		}

		return min
	}
}

func (list *LinkedList) String() string {

	var cNode = list.head
	var str = "HEAD -> " + cNode.String() + "\nTAIL: " + list.tail.String() + "\n"

	for cNode != nil {
		str += cNode.String() + "\n"
		cNode = cNode.next
	}
	return str
}
