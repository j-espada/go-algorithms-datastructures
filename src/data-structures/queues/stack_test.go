package queues

import (
	"math/rand"
	"testing"
)

const MaxValue = 20

func generateData(n int) []int {
	var data = make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = rand.Int() % MaxValue
	}
	return data
}

func populateStack(stack *Stack, data []int) {
	for i := 0; i < len(data); i++ {
		stack.Push(data[i])
	}
}

func TestStack_CreateStackWithDefinedCapacity(t *testing.T) {
	stackCapacity := 10
	stack := CreateStack(stackCapacity)

	if stack.capacity != stackCapacity {
		t.Error("Expected: ", stackCapacity, "got: ", stack.capacity)
	}

	if stack.top != 0 {
		t.Error("Expected: 0", stack.capacity, "got: ", stack.top)
	}

	if stack.IsEmpty() != true {
		t.Error("Expected: true", "got: false")
	}
}

func TestStack_CreateStackWithDefaultCapacity(t *testing.T) {
	var stack = CreateStackDefaultCapacity()

	if stack.capacity != DefaultSize {
		t.Error("Expected: ", DefaultSize, "got: ", stack.capacity)
	}

	if stack.top != 0 {
		t.Error("Expected: 0", stack.capacity, "got: ", stack.top)
	}

	if stack.IsEmpty() != true {
		t.Error("Expected: true", "got: false")
	}
}

func TestStack_Push(t *testing.T) {
	var n = 10
	var data = generateData(n)
	var stack = CreateStack(n)
	populateStack(stack, data)

	for i := 0; i < n; i++ {
		if stack.stack[i] != data[i] {
			t.Error("Expected: ", data[i], "got: ", stack.stack[i])
		}
	}

	if stack.top != n {
		t.Error("Expected: ", n, "got: ", stack.top)
	}
}

func TestStack_IsEmpty(t *testing.T) {
	var n = 10
	var stack = CreateStack(n)

	if stack.IsEmpty() != true {
		t.Error("Expected: ", true, "got: ", false)
	}
}

func TestStack_IsEmptyMultiplePop(t *testing.T) {
	var n = 10
	var data = generateData(n)
	var stack = CreateStack(n + 1)
	populateStack(stack, data)

	for i := n; i >= 2; i-- {
		stack.Pop()
		if false != stack.IsEmpty() {
			t.Error("Expected: ", false, " got: ", stack.IsEmpty())
		}
	}

	stack.Pop()

	if true != stack.IsEmpty() {
		t.Error("Expected: ", true, " got: ", stack.IsEmpty())
	}

}

func TestStack_Pop(t *testing.T) {
	var n = 10
	var data = generateData(n)
	var stack = CreateStack(n + 1)
	populateStack(stack, data)

	for stack.IsEmpty() == false {
		stack.Pop()
	}

	if stack.IsEmpty() != true {
		t.Error("Expected: ", true, "got: ", false)
	}
}

func TestStack_Top(t *testing.T) {
	var n = 10
	var data = generateData(n)
	var stack = CreateStack(n + 1)
	populateStack(stack, data)

	if stack.IsEmpty() != false {
		t.Error("Expected: ", false, " got: ", stack.IsEmpty())
	}

	if stack.Top() != data[n-1] {
		t.Error("Expected: ", data[n-1], " got: ", stack.Top())
	}
}

func TestStack_TopEmptyStack(t *testing.T) {
	var stack = CreateStackDefaultCapacity()

	if stack.IsEmpty() != true {
		t.Error("Expected: ", true, " got: ", stack.IsEmpty())
	}

	if stack.Top() != nil {
		t.Error("Expected: ", nil, " got: ", stack.Top())
	}
}

func TestStack_EmptyStackGotNil(t *testing.T) {
	var n = 10
	var stack = CreateStack(n)
	var pop = stack.Pop()
	if pop != nil {
		t.Error("Expected: ", nil, "got: ", pop)
	}
}

func TestStack_DuplicateSize(t *testing.T) {
	var n = 10
	var stack = CreateStack(n)
	var stackCap = stack.capacity
	stack.duplicateSizeAndCopy()

	if stack.capacity != stackCap*2 {
		t.Error("Expected: ", stack.capacity, "got: ", stackCap*2)
	}
}
