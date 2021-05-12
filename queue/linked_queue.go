package queue

import "github.com/zhao2490/algorithms/LinkedList"

/*
基于链表实现队列
*/

type LinkedQueue struct {
	head *LinkedList.SingleLinkedNode
	tail *LinkedList.SingleLinkedNode
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{}
}

func (q *LinkedQueue) enqueue(item int32) bool {
	newNode := &LinkedList.SingleLinkedNode{
		Val:  item,
		Next: q.head,
	}
	if q.tail == q.head {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.Next = newNode
	}
	return true
}

func (q *LinkedQueue) dequeue() int32 {
	if q.head == nil {
		return 0
	}
	tmp := q.head
	q.head = q.head.Next
	return tmp.Val
}
