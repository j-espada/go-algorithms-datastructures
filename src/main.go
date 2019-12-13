package main

import "fmt"
import "./algorithms/sorting"

type Person struct {
	Name string
	Age  int
}

// ByAge implements sort.Interface based on the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }


func main() {
	arr:= []int{9,8,7,6,5,4,3,2,1,0}
	fmt.Println(arr)
	var arrayS = sorting.BubbleSort(arr, len(arr))
	fmt.Println(arrayS)

	family := []Person{
		{"Alice", 23},
		{"Eve", 2},
		{"Bob", 25},
	}

	fmt.Println(family)
	var sortedFamily = sorting.BubbleSortGeneric(ByAge(family))
	fmt.Println(sortedFamily)

}