package main

import (
	"./algorithms/sorting"
	"./data-structures/stack"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// ByAge implements sort.Interface based on the Age field.
type ByAge []Person
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type NodeInterfaceImpl struct {
	Value int
}
func (a NodeInterfaceImpl) String() string {
	return fmt.Sprintln(a.Value)
}


func main() {
	stackTest()
}

func stackTest() {
	var s = stack.CreateStack(10)
	fmt.Println(s)
	var x = &NodeInterfaceImpl{Value:10}
	var n = &stack.Node{Value:x}
	s.Push(n)
	fmt.Println(s)

}

func sortTest(){
	arr:= []int{9,8,8,7,6,5,4,3,2,1,0}

	//var arrayS = sorting.CountingSortInteger(arr, len(arr), 9)
	var arrayS = sorting.BucketSortInteger(arr, len(arr),9)
	fmt.Println(arrayS)
	/*
		family := []Person{
			{"Alice", 23},
			{"Eve", 2},
			{"Bob", 25},
		}

		fmt.Println(family)
	*/

	//var sortedFamily = sorting.BubbleSortGeneric(ByAge(family))
	//fmt.Println(sortedFamily)
}