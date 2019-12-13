package main

import "fmt"
import "./algorithms/sorting"

func main() {
	arr:= []int{9,8,7,6,5,4,3,2,1,0}
	fmt.Println(arr)
	var arrayS = sorting.CountingSortInteger(arr, len(arr), 9)
	fmt.Println(arrayS)
}