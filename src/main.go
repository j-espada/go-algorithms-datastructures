package main

import "fmt"
import "./algorithms/sorting"

func main() {
	arr:= []int{9,8,7,6,5,4,3,2,1,0}
	fmt.Println(arr)
	var arrayS = sorting.BubbleSort(arr, len(arr))
	fmt.Println(arrayS)
}