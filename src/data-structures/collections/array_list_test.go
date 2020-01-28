package collections

import (
	"testing"
)

func TestCreateArrayList(t *testing.T) {
	var capacity = 10
	var arrayList = CreateArrayList(capacity)

	if arrayList.Size() != 0 {
		t.Error("Expected size to be: ", 0, " got: ", arrayList.Size())
	}

	if arrayList.Capacity() != capacity {
		t.Error("Expected size to be: ", capacity, " got: ", arrayList.Capacity())
	}
}

func TestArrayList_AddOneItem(t *testing.T) {
	var item = 12
	var arrayList = CreateArrayListDefaultCapacity()
	arrayList.Add(item)

	if arrayList.Contains(item, Equals) != true {
		t.Error("Expected to be: ", true, " got: ", arrayList.Contains(item, Equals))
	}

	if arrayList.size != 1 {
		t.Error("Expected to be: ", 1, " got: ", arrayList.size)
	}

	if arrayList.Get(0) != item {
		t.Error("Expected to be: ", item, " got: ", arrayList.Get(0))
	}
}

func TestArrayList_AddMultipleItem(t *testing.T) {
	var n = 10
	var arrLstCap = 2
	var data = generateData(n)
	var arrayList = CreateArrayList(arrLstCap)
	insertData(arrayList, data)

	for _, item := range data {
		if arrayList.Contains(item, Equals) != true {
			t.Error("Expected to be: ", true, " got: ", arrayList.Contains(item, Equals))

		}
	}

	if arrayList.size != n {
		t.Error("Expected to be: ", 1, " got: ", arrayList.size)
	}

	for i, item := range data {

		k := arrayList.Get(i)

		if k != item {
			t.Error("Expected to be: ", item, " got: ", k)
		}
	}

}

func TestArrayList_AddAll(t *testing.T) {
	var n1 = 10
	var n2 = 8
	var n3 = n1 + n2
	var data1 = generateData(n1)
	var data2 = generateData(n2)
	var arrayList1 = CreateArrayList(n1)
	var arrayList2 = CreateArrayList(n2)

	insertData(arrayList1, data1)
	insertData(arrayList2, data2)
	arrayList1.AddAll(arrayList2)

	if arrayList1.size != n3 {
		t.Error("Expected to be: ", n3, " got: ", arrayList1.size)
	}

	for _, item := range data1 {
		if arrayList1.Contains(item, Equals) != true {
			t.Error("Expected to be: ", true, " got: ", arrayList1.Contains(item, Equals))

		}
	}

	for _, item := range data2 {
		if arrayList1.Contains(item, Equals) != true {
			t.Error("Expected to be: ", true, " got: ", arrayList1.Contains(item, Equals))
		}
	}
}

func TestArrayList_Clear(t *testing.T) {
	var n = 10
	var data = generateData(n)
	var arrayList = CreateArrayList(n)
	insertData(arrayList, data)

	arrayList.Clear()

	for _, item := range data {
		if arrayList.Contains(item, Equals) != false {
			t.Error("Expected to be: ", false, " got: ", arrayList.Contains(item, Equals))

		}
	}
}

func TestArrayList_Get(t *testing.T) {
	var dataToInsert = []int{10, 2, 55, 10}
	var arrayList = CreateArrayList(len(dataToInsert))
	insertData(arrayList, dataToInsert)
	var get = arrayList.Get(0)

	if get != dataToInsert[0] {
		t.Error("Expected to be: ", dataToInsert[0], " got: ", get)
	}
}

func TestArrayList_GetNegativePosition(t *testing.T) {
	var dataToInsert = []int{10, 2, 55, 10}
	var arrayList = CreateArrayList(len(dataToInsert))
	insertData(arrayList, dataToInsert)
	var get = arrayList.Get(-1)

	if get != nil {
		t.Error("Expected to be: ", nil, " got: ", get)
	}
}

func TestArrayList_GetLargerThanLst(t *testing.T) {
	var dataToInsert = []int{10, 2, 55, 10}
	var arrayList = CreateArrayList(len(dataToInsert))
	insertData(arrayList, dataToInsert)
	var get = arrayList.Get(len(dataToInsert) + 1)

	if get != nil {
		t.Error("Expected to be: ", nil, " got: ", get)
	}
}

func TestArrayList_IndexOf(t *testing.T) {
	var dataToInsert = []int{10, 2, 55, 10}
	var arrayList = CreateArrayList(len(dataToInsert))
	insertData(arrayList, dataToInsert)

	var val = 10
	var pos = 0
	var index = arrayList.IndexOf(val, Equals)

	if index != pos {
		t.Error("Expected to be: ", pos, " got: ", index)
	}
}

func TestArrayList_LastIndexOf(t *testing.T) {
	var dataToInsert = []int{10, 2, 55, 10}
	var arrayList = CreateArrayList(len(dataToInsert))
	insertData(arrayList, dataToInsert)

	var val = 10
	var pos = 3
	var index = arrayList.LastIndexOf(val, Equals)

	if index != pos {
		t.Error("Expected to be: ", pos, " got: ", index)
	}
}

func TestArrayList_LastIndexOfNotExists(t *testing.T) {
	var dataToInsert = []int{10, 2, 55, 10}
	var arrayList = CreateArrayList(len(dataToInsert))
	insertData(arrayList, dataToInsert)

	var val = 102
	var pos = -1
	var index = arrayList.LastIndexOf(val, Equals)

	if index != pos {
		t.Error("Expected to be: ", pos, " got: ", index)
	}
}

func TestArrayList_IndexOfNoElm(t *testing.T) {
	var dataToInsert = []int{10, 2, 55, 10}
	var arrayList = CreateArrayList(len(dataToInsert))
	insertData(arrayList, dataToInsert)

	var val = 110
	var pos = -1
	var index = arrayList.IndexOf(val, Equals)

	if index != pos {
		t.Error("Expected to be: ", pos, " got: ", index)
	}
}

func TestArrayList_ToArray(t *testing.T) {
	var dataToInsert = []int{10, 2, 55, 10}
	var arrayList = CreateArrayList(len(dataToInsert))
	insertData(arrayList, dataToInsert)
	var toArray = arrayList.ToArray()

	if len(toArray) != len(dataToInsert) {
		t.Error("Expected to be: ", len(dataToInsert), " got: ", len(toArray))
	}
}

func TestArrayList_Max(t *testing.T) {
	var arrayList = CreateArrayList(4)
	var maxValEmpty = arrayList.Max(Comparable)

	if nil != maxValEmpty {
		t.Error("Expected: ", nil, " got: ", maxValEmpty)
	}

	var dataToInsert = []int{10, 2, 55, 10}
	insertData(arrayList, dataToInsert)
	var max = arrayList.Max(Comparable)
	var maxValueData = 55

	if max != maxValueData {
		t.Error("Expected: ", maxValueData, " got: ", max)
	}
}

func TestArrayList_Min(t *testing.T) {
	var arrayList = CreateArrayList(4)
	var minValEmpty = arrayList.Min(Comparable)

	if nil != minValEmpty {
		t.Error("Expected: ", nil, " got: ", minValEmpty)
	}

	var dataToInsert = []int{10, 2, 55, 10}
	insertData(arrayList, dataToInsert)
	var min = arrayList.Min(Comparable)
	var minValueData = 2

	if min != minValueData {
		t.Error("Expected: ", minValueData, " got: ", min)
	}
}

func TestArrayList_RemoveFirstElm(t *testing.T) {
	var arrayList = CreateArrayList(4)
	var dataToInsert = []int{10, 2, 55, 55}
	insertData(arrayList, dataToInsert)

	var deletePos = 0
	var item = dataToInsert[deletePos]
	arrayList.Remove(item, Equals)
	var contains = arrayList.Contains(item, Equals)

	if contains != false {
		t.Error("Expected: ", false, " got: ", contains)
	}

	var dataToInsertRemoved = []int{2, 55, 55}
	removeUtil(dataToInsertRemoved, arrayList.arr, 3, t)
}

func TestArrayList_RemoveLastElm(t *testing.T) {
	var arrayList = CreateArrayList(4)
	var dataToInsert = []int{10, 2, 155, 55}
	insertData(arrayList, dataToInsert)

	var deletePos = 3
	var item = dataToInsert[deletePos]
	arrayList.Remove(item, Equals)
	var contains = arrayList.Contains(item, Equals)

	if contains != false {
		t.Error("Expected: ", false, " got: ", contains)
	}

	var dataToInsertRemoved = []int{10, 2, 155}
	removeUtil(dataToInsertRemoved, arrayList.arr, 3, t)
}

func TestArrayList_RemoveMiddleElm(t *testing.T) {
	var arrayList = CreateArrayList(4)
	var dataToInsert = []int{10, 2, 55, 55}
	insertData(arrayList, dataToInsert)

	var deletePos = 1
	var item = dataToInsert[deletePos]
	arrayList.Remove(item, Equals)
	var contains bool = arrayList.Contains(item, Equals)

	if contains != false {
		t.Error("Expected: ", false, " got: ", contains)
	}

	var dataToInsertRemoved = []int{10, 55, 55}
	removeUtil(dataToInsertRemoved, arrayList.arr, 3, t)
}

func removeUtil(expectedArr []int, newArr []interface{}, n int, t *testing.T) {
	for i := 0; i < n; i++ {
		if Equals(expectedArr[i], newArr[i]) != true {
			t.Error("Expected: ", expectedArr[i], " got: ", newArr[i])
		}
	}
}

func TestArrayList_SubList(t *testing.T) {
	var arrayList *ArrayList = CreateArrayList(6)

	var dataToInsert = []int{1, 2, 3, 4, 5, 6}
	insertData(arrayList, dataToInsert)

	var sbl = arrayList.SubList(-1, 3)

	if nil != sbl {
		t.Error("Expected: ", nil, " got: ", sbl)
	}

	sbl = arrayList.SubList(0, -3)

	if nil != sbl {
		t.Error("Expected: ", nil, " got: ", sbl)
	}

	sbl = arrayList.SubList(6, 5)

	if nil != sbl {
		t.Error("Expected: ", nil, " got: ", sbl)
	}

	var fromIndex = 0
	var toIndex = 3
	var subListData = []int{1, 2, 3}
	var expectedSbl = CreateArrayList(4)
	insertData(expectedSbl, subListData)

	arrListSbl := arrayList.SubList(fromIndex, toIndex)

	if arrListSbl.Size() != len(subListData) {
		t.Error("Expected: ", len(subListData), " got: ", arrListSbl.Size())
	}
}
