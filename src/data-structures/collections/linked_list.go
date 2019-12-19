package collections

type Node struct {
	Value CollectionItem
	next *Node
}

func createNode(item CollectionItem) *Node{
  return &Node{
	  item,
	  nil,
  }
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func createLinkedList() *LinkedList  {
	return &LinkedList{
		nil,
		nil,
		0,
	}
}

// Appends the specified element to the end of this list
func (list *LinkedList) Add(item CollectionItem) {

	var node = createNode(item)
	var currentNode = list.head

	for currentNode.next != nil{
		currentNode = currentNode.next
	}

	currentNode.next = node
}

// Appends all of the elements in the specified collection to the end of this list, in the order that they are returned by the specified collection's iterator
func (list *LinkedList) AddAll(collectionInterface CollectionInterface) {

}

// Returns the index of the first occurrence of the specified element in this list, or -1 if this list does not contain the element.
func (list *LinkedList) IndexOf(item CollectionItem) int {
	var i = 0
	var currentNode = list.head
	var elm CollectionItem = nil

	for currentNode.next != nil {

		if item.Equals(currentNode.Value){
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

	var i = 0
	var currentNode = list.head
	var elm CollectionItem = nil

	for currentNode.next != nil {

		if item.Equals(currentNode.Value){
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

//Removes all of the elements from this list
func (list *LinkedList) Clear() { list.head.next = nil}

//Returns true if this list contains the specified element.
func (list *LinkedList) Contains(item CollectionItem) bool {
	var currentNode = list.head
	for currentNode.next != nil {
		if item.Equals(currentNode.Value){
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// Returns the number of elements in this list.
func (list *LinkedList) Size() int{
	return list.size
}
// Removes the first occurrence of the specified element from this list, if it is present
func (list *LinkedList) Remove(item CollectionItem) {}

// Returns the element at the specified position in this list.
func (list *LinkedList) Get(i int) CollectionItem {

	if i > list.size || i < 0 {
		return nil
	}

	var currentNode = list.head
	var k = 0
	var item CollectionItem = nil
	for currentNode.next != nil {
		if i == k  {
			item = currentNode.Value
			break
		}
		currentNode = currentNode.next
		k++
	}

	return item
}

// Returns a view of the portion of this list between the specified fromIndex, inclusive, and toIndex, exclusive.
func (list *LinkedList) SubList(fromIndex int, toIndex int) CollectionInterface {

	if fromIndex < 0 || fromIndex > toIndex || fromIndex > list.size{
		return nil
	}
	if toIndex < fromIndex || toIndex > list.size {
		return nil
	}

	var subList = createLinkedList()
	var i = 0
	var currentNode = list.head

	for currentNode.next != nil  {
		if i >= fromIndex && i <= toIndex {
			subList.Add(currentNode.Value)
		}
		i++
	}

	return subList
}

// Returns an array containing all of the elements in this list in proper sequence (from first to last element)
func (list *LinkedList) ToArray() []CollectionItem {
	var arr = make([]CollectionItem, list.size)
	var k = 0
	var currentNode = list.head.next
	for currentNode.next != nil {
		arr[k] = currentNode.Value
		currentNode = currentNode.next
		k++
	}
	return arr
}