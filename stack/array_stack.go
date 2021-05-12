package stack

type ArrayStack struct {
	items []string // 数组
	count int      // 栈中元素个数
	n     int      // 栈的大小
}

func NewArrayStack(n int) *ArrayStack {
	return &ArrayStack{
		items: make([]string, n, n),
		count: n,
		n:     n,
	}
}

// 入栈操作
func (s *ArrayStack) Push(item string) bool {
	// 数组空间不够了，直接返回false，入栈失败
	if s.count == s.n {
		return false
	}
	// 将item放入count的位置，并且count加一
	s.items[s.count] = item
	s.count++
	return true
}

// 出栈操作
func (s *ArrayStack) Pop() (bool, string) {
	// 栈为空，则直接返回nil
	if s.count == 0 {
		return false, ""
	}
	// 返回下标为count-1的数组元素，并且栈中元素个数count减一
	tmp := s.items[s.count-1]
	s.count--
	return true, tmp
}


