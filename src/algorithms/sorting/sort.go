package sorting

import "fmt"
type Interface interface {
	 	// Len is the number of elements in the collection.
	 	Len() int
	  	// Less reports whether the element with
	  	// index i should sort before the element with index j.
	  	Less(i, j int) bool
	 	// Swap swaps the elements with indexes i and j.
	  	Swap(i, j int)
	 	// Find universe size
	 	UniverseSize() int
}

func BubbleSortGeneric(A Interface) Interface {
	var n = A.Len()
	var N = n - 1
	for i := 0; i < N; i++ {
		for j := N; j >= i + 1 ; j-- {
			var k = j - 1
			if A.Less(j,k) {
				A.Swap(j,k)
			}
		}
	}
	return A
}

func InsertionSortGeneric(A Interface) Interface {
	var n = A.Len()
	var N = n - 1
	for j := 0; j <= N; j++  {
		var smallest = j
		for i:= j + 1; i <= N; i++ {
			if A.Less(i,smallest){
				smallest = i
			}
			A.Swap(j,i)
		}
	}
	return A
}

func BubbleSortInteger(A []int, n int) []int {
	var N = n - 1
	for i := 0; i < N; i++ {
		for j := N; j >= i + 1 ; j-- {
			var k = j - 1
			if A[j] < A[k] {
				exchange(A, k, j)
			}
		}
	}
	return A
}

func InsertionSortInteger(A []int, n int) []int {
	var N = n - 1
	for j := 0; j <= N; j++  {
		var smallest = j
		for i:= j + 1; i <= N; i++ {
			if A[i] < A[smallest] {
				smallest = i
			}
			exchange(A, j , i)
		}
	}
	return A
}

func CountingSortInteger(A []int, n int, k int) []int {
	var B = make([]int, n)
	var C = make([]int, k+1) // counting array
	fmt.Println(C)
	// counting
	for j:= 1; j < n ; j++  {
		C[A[j]] = C[A[j]] + 1
	}
	fmt.Println(C)
	for i:= 1; i < k ; i++ {
		C[i] = C[i] + C[i - 1]
	}
	fmt.Println(C)
	for j:= n - 1; j >= 1; j-- {
		B[C[A[j]]] = A[j]
		C[A[j]] = C[A[j]] - 1
	}

	return B
}

func exchange(array []int, i int, j int){
	var temp = array[j]
	array[j] = array[i]
	array[i] = temp
}