package collections

import (
	"math/rand"
)

const MaxValue = 20

func generateData(n int) []int {
	var data = make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = rand.Int() % MaxValue
	}
	return data
}

func insertData(lst CollectionInterface, data []int) {
	for _, data := range data {
		lst.Add(data)
	}
}

func Equals(o1 interface{}, o2 interface{}) bool {
	oo1, ok := o1.(int)
	oo2, ok2 := o2.(int)
	if ok != true || ok2 != true {
		panic("Invalid cast")
	}
	return oo1 == oo2
}

func Comparable(o1 interface{}, o2 interface{}) int {
	oo1, ok := o1.(int)
	oo2, ok2 := o2.(int)
	if ok != true || ok2 != true {
		panic("Invalid cast")
	}
	return oo1 - oo2
}
