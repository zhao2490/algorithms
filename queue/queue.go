package queue

type ArrayQueue struct {
	items []string
	n     int
	head  int // 表示队头下标
	tail  int // 表示队尾下标
}

func NewArrayQueue(capacity int) *ArrayQueue {
	return &ArrayQueue{
		items: make([]string, capacity, capacity),
		n:     capacity,
	}
}

func (q *ArrayQueue) enqueue(item string) bool {
	// 表示队尾没有空间了
	if q.tail == q.n {
		if q.head == 0 {
			return false
		}
		// 数据搬迁
		for i := q.head; i < q.tail; i++ {
			q.items[i-q.head] = q.items[i]
		}
		// 搬迁完重新更新head和tail
		q.tail -= q.head
		q.head = 0
	}
	q.items[q.tail] = item
	q.tail++
	return true
}

func (q *ArrayQueue) dequeue() string {
	// head == tail 表示队列为空
	if q.head == q.tail {
		return ""
	}
	ret := q.items[q.head]
	q.head++
	return ret
}