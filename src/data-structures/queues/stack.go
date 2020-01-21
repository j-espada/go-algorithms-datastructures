package queues

import "fmt"

type Stack struct {
	capacity int
	top      int
	stack    []interface{}
}

func CreateStack(capacity int) *Stack {
	return &Stack{capacity: capacity,
		stack: make([]interface{}, capacity),
		top:   0}
}

func CreateStackDefaultCapacity() *Stack {
	return &Stack{capacity: DefaultSize,
		stack: make([]interface{}, DefaultSize),
		top:   0}
}

func (stack *Stack) IsEmpty() bool {
	return stack.top == 0
}

func (stack *Stack) Push(node interface{}) {

	if stack.top == stack.capacity-1 {
		stack.duplicateSizeAndCopy()
	}

	stack.stack[stack.top] = node
	stack.top = stack.top + 1
}

func (stack *Stack) duplicateSizeAndCopy() {
	var newCap = stack.capacity * 2
	var newStackArray = make([]interface{}, newCap)
	for i := 0; i < stack.capacity; i++ {
		newStackArray[i] = stack.stack[i]
	}
	stack.capacity = newCap
	stack.stack = newStackArray
}

func (stack *Stack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	} else {
		var toReturn = stack.stack[stack.top-1]
		stack.stack[stack.top-1] = nil
		stack.top = stack.top - 1
		return toReturn
	}
}

func (stack *Stack) Top() interface{} {
	if stack.IsEmpty() {
		return nil
	} else {
		return stack.stack[stack.top-1]
	}
}

func (stack *Stack) String() string {
	var str = ""
	str += fmt.Sprintln("Capacity: ", stack.capacity)
	str += fmt.Sprintln("Top: ", stack.top)
	str += fmt.Sprint(stack.stack)
	return str
}
