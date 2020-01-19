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
