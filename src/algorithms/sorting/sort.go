package sorting


func BubbleSort(array []int, n int) []int {

	var arraySize = n - 1

	for i := 0; i < arraySize; i++ {
		for j := arraySize; j >= i + 1 ; j-- {
			var k = j - 1
			if array[j] < array[k] {
				//var temp = array[k]
				exchange(array, k, j)
				//array[j - 1] = array[j]
				//array[j] = temp
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