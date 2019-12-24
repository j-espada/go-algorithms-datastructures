package collections

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
)

const MaxValue = 20

type T struct {
	Value int
}

func (item T) Equals(other CollectionItem) bool {
	otherT, ok := other.(T)
	if !ok {
		return false
	}
	return otherT.Value == item.Value
}

func (item T) Compare(other CollectionItem) (int, error) {
	otherT, ok := other.(T)
	if !ok {
		return -1, errors.New("IMPOSSIBLE TO COMPARE DIFFERENT TYPES")
	}
	return otherT.Value - item.Value, nil
}

func (item T) String() string {
	return fmt.Sprint(item.Value)
}

func generateData(n int) [] T {
	var data = make([] T, n)
	for i := 0; i < n  ; i++  {
		data[i] = T{Value:rand.Int() % MaxValue}
	}
	return data
}

func insertData(linkedList *LinkedList, data []T) {
	for _, data := range data {
		linkedList.Add(data)
	}
}

func TestCreateNode(t *testing.T)  {
	var val = T{Value: 10}
	var node = createNode(val)
	if node.Value != val {
		t.Error("Expected: ", val, "got: ", node.Value)
	}
}

func TestCreateLinkedList(t *testing.T) {
	var linkedList = CreateLinkedList()
	var expectedSize = 0

	if  expectedSize != linkedList.size{
		t.Error("Expected: ", expectedSize, "got: ", linkedList.size)
	}

	if linkedList.head == nil {
		t.Error("Expected: ", expectedSize, "got: ", linkedList.size)
	}

	if linkedList.head.Value != nil {
		t.Error("Expected: ", nil, "got: ", linkedList.head.Value)
	}
}


func TestLinkedList_Add(t *testing.T) {
	var linkedList = CreateLinkedList()
	var n = 100
	var dataToInsert = generateData(n)

	for _, data := range dataToInsert {
		linkedList.Add(data)
	}

	if linkedList.size != n {
		t.Error("Expected: ", n, "got: ", linkedList.size)
	}

	var k = 0
	var currentNode = linkedList.head.next
	for currentNode != nil  {
		if ! currentNode.Value.Equals(dataToInsert[k]) {
			t.Error("Expected: ", dataToInsert[k], "got: ", currentNode.Value)
		}
		currentNode = currentNode.next
		k++
	}

	if k != n {
		t.Error("Expected: ", n, "got: ", k)

	}
}

func TestLinkedList_AddAll(t *testing.T) {

}

func TestLinkedList_IndexOf(t *testing.T) {

	var linkedList = CreateLinkedList()
	var dataToInsert = []T{{Value: 10}, {Value: 2}, {Value: 55}, {Value: 12}}

	insertData(linkedList, dataToInsert)
	var pos = 0
	var elm = T{Value: 10}
	var posLst = linkedList.IndexOf(elm)

	if pos != posLst{
		t.Error("Expected: ", pos, "got: ", posLst)
	}
}

func TestLinkedList_IndexOf_RepeatElm(t *testing.T) {

	var linkedList = CreateLinkedList()
	var dataToInsert = []T{{Value: 10}, {Value: 2}, {Value: 55}, {Value: 55}}

	insertData(linkedList, dataToInsert)
	var pos = 2
	var elm = T{Value: 55}
	var posLst = linkedList.IndexOf(elm)

	if pos != posLst{
		t.Error("Expected: ", pos, "got: ", posLst)
	}
}

func TestLinkedList_IndexOf_NoElement(t *testing.T) {

	var linkedList = CreateLinkedList()
	var dataToInsert = []T{{Value: 10}, {Value: 2}, {Value: 55}, {Value: 12}}

	insertData(linkedList, dataToInsert)
	var pos = -1
	var elm = T{Value: 100}

	var posLst = linkedList.IndexOf(elm)

	if pos != posLst{
		t.Error("Expected: ", pos, "got: ", posLst)
	}
}

func TestLinkedList_LastIndexOf(t *testing.T) {

	var linkedList = CreateLinkedList()
	var dataToInsert = []T{{Value: 10}, {Value: 2}, {Value: 55}, {Value: 10}}

	insertData(linkedList, dataToInsert)
	var pos = 3
	var elm = T{Value: 10}
	var posLst = linkedList.LastIndexOf(elm)

	if pos != posLst{
		t.Error("Expected: ", pos, "got: ", posLst)
	}
}

func TestLinkedList_Clear(t *testing.T) {
	var linkedList = CreateLinkedList()
	var n = 10
	var dataToInsert = generateData(n)
	insertData(linkedList, dataToInsert)

	if n != linkedList.Size() {
		t.Error("Expected: ", n, " got: ", linkedList.size)
	}

	linkedList.Clear()
	if 0 != linkedList.Size() {
		t.Error("Expected: ", 0, " got: ", linkedList.size)

	}
}

func TestLinkedList_Contains(t *testing.T) {

	var linkedList = CreateLinkedList()
	var dataToInsert = []T{{Value: 10}, {Value: 2}, {Value: 55}, {Value: 10}}
	insertData(linkedList, dataToInsert)

	var contains = linkedList.Contains(T{Value:10})

	if true != contains {
		t.Error("Expected: ", true, " got: ", contains)

	}
}

func TestLinkedList_RemoveFirstElm(t *testing.T) {

	var linkedList = CreateLinkedList()
	var dataToInsert = []T{{Value: 10}, {Value: 2}, {Value: 55}, {Value: 110}}
	insertData(linkedList, dataToInsert)

	var toBeRemoved = T{Value:10}
	linkedList.Remove(toBeRemoved)
	var contains = linkedList.Contains(toBeRemoved)

	if false != contains {
		t.Error("Expected: ", false, " got: ", contains)
	}
}

func TestLinkedList_RemoveLastElm(t *testing.T) {

	var linkedList = CreateLinkedList()
	var dataToInsert = []T{{Value: 10}, {Value: 2}, {Value: 55}, {Value: 110}}
	insertData(linkedList, dataToInsert)

	var toBeRemoved = T{Value:110}
	linkedList.Remove(toBeRemoved)
	var contains = linkedList.Contains(toBeRemoved)

	if false != contains {
		t.Error("Expected: ", false, " got: ", contains)
	}
}

func TestLinkedList_RemoveMiddleElm(t *testing.T) {

	var linkedList = CreateLinkedList()
	var dataToInsert = []T{{Value: 10}, {Value: 2}, {Value: 55}, {Value: 110}}
	insertData(linkedList, dataToInsert)

	var toBeRemoved = T{Value:2}
	linkedList.Remove(toBeRemoved)
	var contains = linkedList.Contains(toBeRemoved)

	if false != contains {
		t.Error("Expected: ", false, " got: ", contains)
	}
}