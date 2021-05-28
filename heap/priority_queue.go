package heap

/*
使用堆实现优先级队列

不能保证相同优先级出队顺序和入队顺序一致
*/

// Node 队列节点
type Node struct {
	value    int
	priority int
}

// PQueue priority queue
type PQueue struct {
	heap []Node

	capacity int
	used     int
}

func adjustHeap(src []Node, start, end int) {
	if start > end {
		return
	}

	// 保证优先级最高的节点在 src[1]的位置
	for i := end / 2; i >= start; i-- {
		high := i
		if src[high].priority < src[2*i].priority {
			high = 2 * i
		}
		if 2*i <= end && src[high].priority < src[2*i+1].priority {
			high = 2*i + 1
		}
		if high == i {
			continue
		}
		src[high], src[i] = src[i], src[high]
	}
}

func NewPriorityQueue(capacity int) PQueue {
	return PQueue{
		heap:     make([]Node, capacity+1, capacity+1),
		capacity: capacity,
		used:     0,
	}
}

func (q *PQueue) Push(node Node) {
	if q.used > q.capacity {
		// 队列已满
		return
	}
	q.used++
	q.heap[q.used] = node
	// 堆化放在 Pop 中
	// adjustHeap(q.heap, 1, q.used)
}

// Pop 出队列
func (q *PQueue) Pop() Node {
	if q.used == 0 {
		return Node{-1, -1}
	}
	// 先堆化 再取结果
	adjustHeap(q.heap, 1, q.used)
	node := q.heap[1]
	q.heap[1] = q.heap[q.used]
	q.used--
	return node
}

// 取队顶元素
func (q *PQueue) Top() Node {
	if q.used == 0 {
		return Node{-1, -1}
	}
	adjustHeap(q.heap, 1, q.used)
	return q.heap[1]
}
