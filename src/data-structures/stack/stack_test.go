package stack

import (
	"fmt"
	"math/rand"
	"testing"
)

const MaxValue = 20

type NodeInterfaceImpl struct {
	Value int
}
func (a NodeInterfaceImpl) String() string {
	return fmt.Sprintln(a.Value)
}

func generateData(n int) [] *Node {
	var data = make([] *Node, n)
	for i := 0; i < n  ; i++  {
		data[i] = &Node{Value:&NodeInterfaceImpl{Value:rand.Int() % MaxValue} }
	}
	return data
}

func populateStack(stack *Stack, data [] *Node) {
	for i:= 0; i < len(data); i++  {
		stack.Push(data[i])
	}
}

func TestNode_String(t *testing.T) {
	var value = 10
	var stringOfValue = fmt.Sprintln(value)
	var node = Node{Value: &NodeInterfaceImpl{Value: value}}
	var nodeStr = node.String()
	if stringOfValue != node.String() {
		t.Error("Expected: ", stringOfValue, "got: ", nodeStr)

	}
}

func TestStack_CreateStackWithDefinedCapacity(t *testing.T) {
	stackCapacity := 10
	stack := CreateStack(stackCapacity)

	if stack.capacity != stackCapacity {
		t.Error("Expected: ", stackCapacity, "got: ",stack.capacity)
	}

	if stack.top != 0 {
		t.Error("Expected: 0", stack.capacity, "got: ", stack.top)
	}

	if stack.IsEmpty() != true {
		t.Error("Expected: true", "got: false")
	}
}

func TestStack_CreateStackWithDefaultCapacity(t *testing.T) {
	stack := CreateStackDefaultCapacity()

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
	for i:= 0; i < n ;i ++  {
		if stack.stack[i] != data[i] {
			t.Error("Expected: ", data[i], "got: ", stack.stack[i])
		}

		if stack.top != i {

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
		t.Error("Expected: ", true , "got: ", false)
	}

}

func TestStack_Pop(t *testing.T) {
	var n = 10
	var data = generateData(n)
	var stack = CreateStack(n+1)
	populateStack(stack, data)

	for stack.IsEmpty() == false {
		stack.Pop()
	}

	if stack.IsEmpty() != true {
		t.Error("Expected: ", true , "got: ", false)

	}
}

func TestStack_EmptyStackGotNil(t *testing.T) {
	var n = 10
	var stack = CreateStack(n)
	var pop = stack.Pop()
	if pop != nil {
		t.Error("Expected: ", nil , "got: ", pop)
	}
}

func TestStack_DuplicateSize(t *testing.T) {

	var n = 10
	var stack = CreateStack(n)
	var stackCap = stack.capacity
	stack.duplicateSizeAndCopy()

	if stack.capacity != stackCap*2 {
		t.Error("Expected: ", stack.capacity , "got: ", stackCap*2)
	}
}