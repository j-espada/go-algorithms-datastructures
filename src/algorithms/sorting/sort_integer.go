package sorting

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
	// counting array make instantiate counting array with zeroes
	var C = make([]int, k + 1)

	// computing frequencies
	for j := 0; j < n ; j++  {
		C[A[j]] = C[A[j]] + 1
	}

	// computing cumulative distribution
	for i := 1; i <= k; i++ {
		C[i] = C[i] + C[i - 1]
	}
	// inserting the elements of A in the correct position in B
	for j := n - 1; j >= 0; j-- {
		B[C[A[j]]-1] = A[j]
		C[A[j]] = C[A[j]] - 1
	}

	return B
}

func BucketSortInteger(A []int, n int, M int) []int {

	var B = make([][]int, n)
	var C = make([]int,n)
	M = M + 1

	for i := 0; i < n ; i++  {
		p := n * A[i] / M
		B[p] = append(B[p], A[i])
	}

	for i := 0; i < n ; i++  {
		B[i] = InsertionSortInteger(B[i], len(B[i]))
	}

	var c = 0
	for i := 0; i < n ; i++  {
		for k := 0; k < len(B[i]); k++ {
			C[c] = B[i][k]
			c++
		}
	}
	return C
}

func exchange(array []int, i int, j int){
	var temp = array[j]
	array[j] = array[i]
	array[i] = temp
}