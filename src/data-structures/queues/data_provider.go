package queues

import "math/rand"

const MaxValue = 20

func GenerateData(n int) []int {
	var data = make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = rand.Int() % MaxValue
	}
	return data
}

func PopulateStack(stack *Stack, data []int) {
	for i := 0; i < len(data); i++ {
		stack.Push(data[i])
	}
}
