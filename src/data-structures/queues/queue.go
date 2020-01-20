package queues

type Queue struct {
	front    int
	size     int
	capacity int
	arr      []interface{}
}

func (queue *Queue) CreateQueue(n int) *Queue {
	return &Queue{front: 0,
		size:     0,
		capacity: n,
		arr:      make([]interface{}, n)}
}

// Check if queue is empty
func (queue *Queue) IsEmpty() bool {
	return queue.size == 0
}

// Push a node into the queue
func (queue *Queue) Push(node interface{}) {
	queue.arr[queue.size] = node
	queue.size++
}

// Pop node from queue
func (queue *Queue) Pop() (node interface{}) {
	res := queue.arr[queue.front]
	queue.front++
	queue.size--
	return res
}

// Pop node from queue
func (queue *Queue) Top() (node interface{}) {
	res := queue.arr[queue.front]
	return res
}
