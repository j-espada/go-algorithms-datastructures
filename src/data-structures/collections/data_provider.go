package collections

import (
	"fmt"
	"math/rand"
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

func (item T) Compare(other CollectionItem) int {
	otherT, ok := other.(T)
	if !ok {
		panic("The two types must be the same")
	}
	return otherT.Value - item.Value
}

func (item T) String() string {
	return fmt.Sprint(item.Value)
}

func generateData(n int) []T {
	var data = make([]T, n)
	for i := 0; i < n; i++ {
		data[i] = T{Value: rand.Int() % MaxValue}
	}
	return data
}

func insertData(lst CollectionInterface, data []T) {
	for _, data := range data {
		lst.Add(data)
	}
}
