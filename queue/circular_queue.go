package queue

import "fmt"

type CircularQueue struct {
	items    []interface{}
	capacity int
	head     int
	tail     int
}

func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		items:    make([]interface{}, capacity),
		capacity: capacity,
	}
}

// 队空条件：head == tail
func (q *CircularQueue) IsEmpty() bool {
	if q.head == q.tail {
		return true
	}
	return false
}

// 队满条件：(tail+1)%capacity == tail
func (q *CircularQueue) IsFull() bool {
	if q.head == (q.tail+1)%q.capacity {
		return true
	}
	return false
}

func (q *CircularQueue) EnQueue(item interface{}) bool {
	if q.IsFull() {
		return false
	}
	q.items[q.tail] = item
	q.tail = (q.tail + 1) % q.capacity
	return true
}

func (q *CircularQueue) DeQueue() interface{} {
	if q.head == q.tail {
		return nil
	}
	ret := q.items[q.tail]
	q.head = (q.head + 1) % q.capacity
	return ret
}

func (q *CircularQueue) String() string {
	if q.IsEmpty() {
		return "empty queue"
	}
	result := "head"
	var i = q.head
	for true {
		result += fmt.Sprintf("<-%+v", q.items[i])
		i = (i + 1) % q.capacity
		if i == q.tail {
			break
		}
	}
	result += "<-tail"
	return result
}
