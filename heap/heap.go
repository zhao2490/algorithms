package heap

type Heap struct {
	a     []int // 数组，从下标1开始存储数据
	n     int   // 堆可以存储的最大数据个数
	count int   // 堆中已经存储的数据个数
}

//init heap
func NewHeap(capacity int) *Heap {
	heap := &Heap{}
	heap.n = capacity
	heap.a = make([]int, capacity+1)
	heap.count = 0
	return heap
}

// 堆化
// 把新插入的元素放到堆的最后，是不符合堆的特性的，需要进行调整，让其重新满足堆的特性
func (heap *Heap) insert(data int) {
	if heap.count == heap.n {
		// 堆满了
		return
	}

	heap.count++
	heap.a[heap.count] = data

	i := heap.count
	parent := i / 2
	for parent > 0 && heap.a[parent] < heap.a[i] { // 自下而上堆化
		swap(heap.a, parent, i)
		i = parent
		parent = i / 2
	}
}

// 自下而上堆化
func (heap *Heap) removeMax() {
	if heap.count == 0 {
		// 堆中没有数据
		return
	}

	swap(heap.a, 1, heap.count)
	heap.count--

	heapifyUpToDown(heap.a, heap.count)
}

// 堆化
func heapifyUpToDown(a []int, count int) {

	for i := 1; i <= count/2; {

		maxIndex := i
		if a[i] < a[i*2] {
			maxIndex = i * 2
		}

		if i*2+1 <= count && a[maxIndex] < a[i*2+1] {
			maxIndex = i*2 + 1
		}

		if maxIndex == i {
			break
		}

		swap(a, i, maxIndex)
		i = maxIndex
	}

}

func swap(a []int, i int, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}
