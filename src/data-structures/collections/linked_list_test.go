package collections

import (
	"reflect"
	"testing"
)

func TestCreateNode(t *testing.T) {
	var val = T{Value: 10}
	var node = createNode(val)
	if node.Value != val {
		t.Error("Expected: ", val, "got: ", node.Value)
	}
}

func TestCreateLinkedList(t *testing.T) {
	var linkedList = CreateLinkedList()
	var expectedSize = 0

	if expectedSize != linkedList.size {
		t.Error("Expected: ", expectedSize, "got: ", linkedList.size)
	}

	if linkedList.head != nil {
		t.Error("Expected: ", nil, "got: ", linkedList.head)
	}

	if linkedList.tail != nil {
		t.Error("Expected: ", nil, "got: ", linkedList.tail)
	}
}

func TestLinkedList_Add(t *testing.T) {
	var linkedList = CreateLinkedList()
	var n = 5
	var dataToInsert = generateData(n)
	insertData(linkedList, dataToInsert)
	if linkedList.size != n {
		t.Error("Expected: ", n, "got: ", linkedList.size)
	}

	var k = 0
	var currentNode = linkedList.head
	for currentNode != nil {
		if !currentNode.Value.Equals(dataToInsert[k]) {
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
	var linkedList, linkedList2 = CreateLinkedList(), CreateLinkedList()
	var n1, n2 = 10, 5
	var dataToInsert1, dataToInsert2 = generateData(n1), generateData(n2)
	insertData(linkedList, dataToInsert1)
	insertData(linkedList2, dataToInsert2)
	linkedList.AddAll(linkedList2)

	var cNode = linkedList2.head.next
	for cNode != nil {

		if true != linkedList.Contains(cNode.Value) {
			t.Error("Expected to have: ", cNode.Value)
		}
		cNode = cNode.next
	}
}

func TestLinkedList_IndexOf(t *testing.T) {

	var linkedList = CreateLinkedList()
	var dataToInsert = []T{{Value: 10}, {Value: 2}, {Value: 55}, {Value: 12}}

	insertData(linkedList, dataToInsert)
	var pos = 0
	var elm = T{Value: 10}
	var posLst = linkedList.IndexOf(elm)

	if pos != posLst {
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

	if pos != posLst {
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

	if pos != posLst {
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

	if pos != posLst {
		t.Error("Expected: ", pos, "got: ", posLst)
	}
}

func TestLinkedList_LastIndexOfNotExists(t *testing.T) {

	var linkedList = CreateLinkedList()
	var dataToInsert = []T{{Value: 10}, {Value: 2}, {Value: 55}, {Value: 10}}
	insertData(linkedList, dataToInsert)

	var pos = -1
	var elm = T{Value: 1022}
	var posLst = linkedList.LastIndexOf(elm)

	if pos != posLst {
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

	if linkedList.head != nil {
		t.Error("Expected: ", nil, " got: ", linkedList.head)

	}

	if linkedList.tail != nil {
		t.Error("Expected: ", nil, " got: ", linkedList.tail)
	}
}

func TestLinkedList_Contains(t *testing.T) {

	var linkedList = CreateLinkedList()
	var dataToInsert = []T{{Value: 10}, {Value: 2}, {Value: 55}, {Value: 10}}
	insertData(linkedList, dataToInsert)

	var contains = linkedList.Contains(T{Value: 10})

	if true != contains {
		t.Error("Expected: ", true, " got: ", contains)

	}
}

func TestLinkedList_RemoveFirstElm(t *testing.T) {

	var linkedList = CreateLinkedList()
	var dataToInsert = []T{{Value: 10}, {Value: 2}, {Value: 55}, {Value: 110}}
	var newHead = linkedList.GetLinkedListNode(1)
	insertData(linkedList, dataToInsert)

	var toBeRemoved = T{Value: 10}
	linkedList.Remove(toBeRemoved)
	var contains = linkedList.Contains(toBeRemoved)

	if false != contains {
		t.Error("Expected: ", false, " got: ", contains)
	}

	if linkedList.size != len(dataToInsert)-1 {
		t.Error("Expected: ", len(dataToInsert)-1, " got: ", linkedList.size)
	}

	if newHead != linkedList.head {
		t.Error("Expected: ", newHead, " got: ", linkedList.head)
	}
}

func TestLinkedList_RemoveLastElm(t *testing.T) {

	var linkedList = CreateLinkedList()
	var dataToInsert = []T{{Value: 10}, {Value: 2}, {Value: 55}, {Value: 110}}
	insertData(linkedList, dataToInsert)
	var newTail = linkedList.GetLinkedListNode(2)

	var toBeRemoved = T{Value: 110}
	linkedList.Remove(toBeRemoved)
	var contains = linkedList.Contains(toBeRemoved)

	if false != contains {
		t.Error("Expected: ", false, " got: ", contains)
	}

	if linkedList.size != len(dataToInsert)-1 {
		t.Error("Expected: ", len(dataToInsert)-1, " got: ", linkedList.size)
	}

	if newTail != linkedList.tail {
		t.Error("Expected: ", newTail, " got: ", linkedList.tail)
	}
}

func TestLinkedList_RemoveMiddleElm(t *testing.T) {

	var linkedList = CreateLinkedList()
	var dataToInsert = []T{{Value: 10}, {Value: 2}, {Value: 55}, {Value: 110}}
	insertData(linkedList, dataToInsert)

	var toBeRemoved = T{Value: 2}
	linkedList.Remove(toBeRemoved)
	var contains = linkedList.Contains(toBeRemoved)

	if false != contains {
		t.Error("Expected: ", false, " got: ", contains)
	}

	if linkedList.size != len(dataToInsert)-1 {
		t.Error("Expected: ", len(dataToInsert)-1, " got: ", linkedList.size)
	}
}

func TestLinkedList_RemoveNotExistingElm(t *testing.T) {

	var linkedList = CreateLinkedList()
	var dataToInsert = []T{{Value: 10}, {Value: 2}, {Value: 55}, {Value: 110}}
	insertData(linkedList, dataToInsert)

	var toBeRemoved = T{Value: 1}
	linkedList.Remove(toBeRemoved)
	var contains = linkedList.Contains(toBeRemoved)

	if false != contains {
		t.Error("Expected: ", false, " got: ", contains)
	}

	if linkedList.size != len(dataToInsert) {
		t.Error("Expected: ", len(dataToInsert)-1, " got: ", linkedList.size)
	}
}

func TestLinkedList_Get(t *testing.T) {
	var linkedList = CreateLinkedList()
	var n = 10
	var dataToInsert = generateData(n)
	insertData(linkedList, dataToInsert)
	var get = linkedList.Get(-1)

	if nil != linkedList.Get(-1) {
		t.Error("Expected: ", nil, " got: ", get)

	}

	get = linkedList.Get(n + 1)
	if nil != get {
		t.Error("Expected: ", nil, " got: ", get)

	}

	get = linkedList.Get(4)
	if dataToInsert[4] != get {
		t.Error("Expected: ", nil, " got: ", get)

	}

}

func TestLinkedList_SubList(t *testing.T) {
	var linkedList = CreateLinkedList()
	var dataToInsert = []T{{Value: 1}, {Value: 2}, {Value: 3}, {Value: 4}, {Value: 5}, {Value: 6}}
	insertData(linkedList, dataToInsert)

	var sbl = linkedList.SubList(-1, 3)
	if nil != sbl {
		t.Error("Expected: ", nil, " got: ", sbl)
	}

	sbl = linkedList.SubList(0, -3)

	if nil != sbl {
		t.Error("Expected: ", nil, " got: ", sbl)
	}

	sbl = linkedList.SubList(0, 7)

	if nil != sbl {
		t.Error("Expected: ", nil, " got: ", sbl, linkedList.size)
	}

	sbl = linkedList.SubList(6, 5)

	if nil != sbl {
		t.Error("Expected: ", nil, " got: ", sbl)
	}

	var fromIndex = 0
	var toIndex = 3
	var subListData = []T{{Value: 1}, {Value: 2}, {Value: 3}}
	var expectedSbl = CreateLinkedList()
	insertData(expectedSbl, subListData)
	sbl = linkedList.SubList(fromIndex, toIndex)
	var eq = reflect.DeepEqual(sbl, expectedSbl)
	if true != eq {
		t.Error("Expected: ", true, " got: ", eq, sbl, expectedSbl)

	}
}

func TestLinkedList_ToArray(t *testing.T) {
	var linkedList = CreateLinkedList()
	var n = 10
	var data = generateData(n)

	insertData(linkedList, data)
	var arr = linkedList.ToArray()

	if len(arr) != n {
		t.Error("Expected: ", n, " got: ", len(arr))
	}

}

func TestLinkedList_Max(t *testing.T) {
	var linkedList = CreateLinkedList()
	var maxValEmpty = linkedList.Max()

	if nil != maxValEmpty {
		t.Error("Expected: ", nil, " got: ", maxValEmpty)
	}

	var dataToInsert = []T{{Value: 10}, {Value: 2}, {Value: 55}, {Value: 55}}
	insertData(linkedList, dataToInsert)
	var max = linkedList.Max()
	var maxValueData = T{Value: 55}

	if max != maxValueData {
		t.Error("Expected: ", maxValueData, " got: ", max)
	}
}

func TestLinkedList_Min(t *testing.T) {
	var linkedList = CreateLinkedList()
	var minValEmpty = linkedList.Min()

	if nil != minValEmpty {
		t.Error("Expected: ", nil, " got: ", minValEmpty)
	}

	var dataToInsert = []T{{Value: 10}, {Value: 2}, {Value: 55}, {Value: 55}}
	insertData(linkedList, dataToInsert)
	var min = linkedList.Min()
	var minValueData = T{Value: 2}

	if min != minValueData {
		t.Error("Expected: ", minValueData, " got: ", min)
	}
}
