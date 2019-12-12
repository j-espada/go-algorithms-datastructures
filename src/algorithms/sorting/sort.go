package sorting

type Interface interface {
	 	// Len is the number of elements in the collection.
	 	Len() int
	  	// Less reports whether the element with
	  	// index i should sort before the element with index j.
	  	Less(i, j int) bool
	 	// Swap swaps the elements with indexes i and j.
	  	Swap(i, j int)
}

func BubbleSortGeneric(array Interface) Interface {
	var n = array.Len()
	var arraySize = n - 1
	for i := 0; i < arraySize; i++ {
		for j := arraySize; j >= i + 1 ; j-- {
			var k = j - 1
			if array.Less(j,k) {
				array.Swap(j,k)
			}
		}
	}
	return array
}

func BubbleSort(array []int, n int) []int {
	var arraySize = n - 1
	for i := 0; i < arraySize; i++ {
		for j := arraySize; j >= i + 1 ; j-- {
			var k = j - 1
			if array[j] < array[k] {
				exchange(array, k, j)
			}
		}
	}
	return array
}

func exchange(array []int, i int, j int){
	var temp = array[j]
	array[j] = array[i]
	array[i] = temp
}