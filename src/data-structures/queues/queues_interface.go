package queues

const DefaultSize = 10

type QueueInterface interface {
	// Check if queue is empty
	IsEmpty() bool
	// Push a node into the queue
	Push(node interface{})
	// Pop node from queue
	Pop() (node interface{})
	// Peek node from queue
	Top() (node interface{})
}
