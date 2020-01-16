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
	var item CollectionItem = T{Value: 12}
	var arrayList = CreateArrayListDefaultCapacity()
	arrayList.Add(item)

	if arrayList.Contains(item) != true {
		t.Error("Expected to be: ", true, " got: ", arrayList.Contains(item))

	}

	if arrayList.size != 1 {
		t.Error("Expected to be: ", 1, " got: ", arrayList.size)

	}
}

func TestArrayList_AddMultipleItem(t *testing.T) {
	var n = 10
	var arrLstCap = 2
	var data = generateData(n)
	var arrayList = CreateArrayList(arrLstCap)
	insertData(arrayList, data)

	for _, item := range data {
		if arrayList.Contains(item) != true {
			t.Error("Expected to be: ", true, " got: ", arrayList.Contains(item))

		}
	}

	if arrayList.size != n {
		t.Error("Expected to be: ", 1, " got: ", arrayList.size)
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
		if arrayList1.Contains(item) != true {
			t.Error("Expected to be: ", true, " got: ", arrayList1.Contains(item))

		}
	}

	for _, item := range data2 {
		if arrayList1.Contains(item) != true {
			t.Error("Expected to be: ", true, " got: ", arrayList1.Contains(item))
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
		if arrayList.Contains(item) != false {
			t.Error("Expected to be: ", false, " got: ", arrayList.Contains(item))

		}
	}
}

func TestArrayList_Get(t *testing.T) {
	var dataToInsert = []T{{Value: 10}, {Value: 2}, {Value: 55}, {Value: 10}}
	var arrayList = CreateArrayList(len(dataToInsert))
	insertData(arrayList, dataToInsert)
	var get = arrayList.Get(0)

	if get != dataToInsert[0] {
		t.Error("Expected to be: ", dataToInsert[0], " got: ", get)
	}
}

func TestArrayList_GetNegativePosition(t *testing.T) {
	var dataToInsert = []T{{Value: 10}, {Value: 2}, {Value: 55}, {Value: 10}}
	var arrayList = CreateArrayList(len(dataToInsert))
	insertData(arrayList, dataToInsert)
	var get = arrayList.Get(-1)

	if get != nil {
		t.Error("Expected to be: ", nil, " got: ", get)
	}
}

func TestArrayList_GetLargerThanLst(t *testing.T) {
	var dataToInsert = []T{{Value: 10}, {Value: 2}, {Value: 55}, {Value: 10}}
	var arrayList = CreateArrayList(len(dataToInsert))
	insertData(arrayList, dataToInsert)
	var get = arrayList.Get(len(dataToInsert) + 1)

	if get != nil {
		t.Error("Expected to be: ", nil, " got: ", get)
	}
}

func TestArrayList_IndexOf(t *testing.T) {
	var dataToInsert = []T{{Value: 10}, {Value: 2}, {Value: 55}, {Value: 10}}
	var arrayList = CreateArrayList(len(dataToInsert))
	insertData(arrayList, dataToInsert)

	var val = T{Value: 10}
	var pos = 0
	var index = arrayList.IndexOf(val)

	if index != pos {
		t.Error("Expected to be: ", pos, " got: ", index)
	}
}

func TestArrayList_LastIndexOf(t *testing.T) {
	var dataToInsert = []T{{Value: 10}, {Value: 2}, {Value: 55}, {Value: 10}}
	var arrayList = CreateArrayList(len(dataToInsert))
	insertData(arrayList, dataToInsert)

	var val = T{Value: 10}
	var pos = 3
	var index = arrayList.LastIndexOf(val)

	if index != pos {
		t.Error("Expected to be: ", pos, " got: ", index)
	}
}

func TestArrayList_LastIndexOfNotExists(t *testing.T) {
	var dataToInsert = []T{{Value: 10}, {Value: 2}, {Value: 55}, {Value: 10}}
	var arrayList = CreateArrayList(len(dataToInsert))
	insertData(arrayList, dataToInsert)

	var val = T{Value: 102}
	var pos = -1
	var index = arrayList.LastIndexOf(val)

	if index != pos {
		t.Error("Expected to be: ", pos, " got: ", index)
	}
}

func TestArrayList_IndexOfNoElm(t *testing.T) {
	var dataToInsert = []T{{Value: 10}, {Value: 2}, {Value: 55}, {Value: 10}}
	var arrayList = CreateArrayList(len(dataToInsert))
	insertData(arrayList, dataToInsert)

	var val = T{Value: 110}
	var pos = -1
	var index = arrayList.IndexOf(val)

	if index != pos {
		t.Error("Expected to be: ", pos, " got: ", index)
	}
}

func TestArrayList_SubList(t *testing.T) {

}

func TestArrayList_ToArray(t *testing.T) {
	var dataToInsert = []T{{Value: 10}, {Value: 2}, {Value: 55}, {Value: 10}}
	var arrayList = CreateArrayList(len(dataToInsert))
	insertData(arrayList, dataToInsert)
	var toArray = arrayList.ToArray()

	if len(toArray) != len(dataToInsert) {
		t.Error("Expected to be: ", len(dataToInsert), " got: ", len(toArray))
	}
}

func TestArrayList_Max(t *testing.T) {
	var arrayList = CreateArrayList(4)
	var maxValEmpty = arrayList.Max()

	if nil != maxValEmpty {
		t.Error("Expected: ", nil, " got: ", maxValEmpty)
	}

	var dataToInsert = []T{{Value: 10}, {Value: 2}, {Value: 55}, {Value: 55}}
	insertData(arrayList, dataToInsert)
	var max = arrayList.Max()
	var maxValueData = T{Value: 55}

	if max != maxValueData {
		t.Error("Expected: ", maxValueData, " got: ", max)
	}
}

func TestArrayList_Min(t *testing.T) {
	var arrayList = CreateArrayList(4)
	var minValEmpty = arrayList.Min()

	if nil != minValEmpty {
		t.Error("Expected: ", nil, " got: ", minValEmpty)
	}

	var dataToInsert = []T{{Value: 10}, {Value: 2}, {Value: 55}, {Value: 55}}
	insertData(arrayList, dataToInsert)
	var min = arrayList.Min()
	var minValueData = T{Value: 2}

	if min != minValueData {
		t.Error("Expected: ", minValueData, " got: ", min)
	}
}
