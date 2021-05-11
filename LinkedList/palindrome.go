package LinkedList

/*
单向链表如何判断字符串是否回文
*/
type SingleLinkedList struct {
	Length int
	Head   *SingleLinkedNode
}

type SingleLinkedNode struct {
	Val  int32
	Next *SingleLinkedNode
}

func NewSingleLinkedList(str string) *SingleLinkedList {
	head := &SingleLinkedNode{}
	var p *SingleLinkedNode = head
	for _, c := range str {
		node := &SingleLinkedNode{
			Val:  c,
			Next: nil,
		}
		p.Next = node
		p = p.Next
	}
	return &SingleLinkedList{
		Length: len(str),
		Head:   head.Next,
	}
}

/*
开一个栈存放链表前半段
时间复杂度 O(n)
空间复杂度 O(1)
*/
func isPalindrome1(l *SingleLinkedList) bool {
	lLen := l.Length
	if lLen == 0 {
		return false
	}
	if lLen == 1 {
		return true
	}

	s := make([]int32, 0, lLen/2)
	cur := l.Head
	for i := 1; i <= lLen; i++ {
		if lLen%2 != 0 && i == (lLen/2+1) { //如果链表有奇数个节点，中间的直接忽略
			cur = cur.Next
			continue
		}
		if i <= lLen/2 { //前一半入栈
			s = append(s, cur.Val)
		} else { //后一半与前一半进行对比
			if s[lLen-i] != cur.Val {
				return false
			}
		}
		cur = cur.Next
	}

	return true
}

/*
找到链表中间节点，将前半部分转置，再从中间向左右遍历对比
时间复杂度 O(n)
*/
func isPalindrome2(l *SingleLinkedList) bool {
	lLen := l.Length
	if lLen == 0 {
		return false
	}
	if lLen == 1 {
		return true
	}
	var pre *SingleLinkedNode = nil
	cur := l.Head
	step := lLen / 2
	for i := 0; i < step; i++ {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	mid := cur

	var left, right *SingleLinkedNode = pre, nil
	if lLen%2 == 0 {
		right = mid
	} else {
		right = mid.Next
	}
	var isPalindrome bool = true
	for left != nil && right != nil {
		if left.Val != right.Val {
			isPalindrome = false
			break
		}
		left = left.Next
		right = right.Next
	}

	// 复原链表
	cur = pre
	pre = mid
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	l.Head = pre

	return isPalindrome
}

/*
快慢指针，将前半部分转置，再从中间向左右遍历对比
时间复杂度 O(n)
*/
func isPalindrome3(l *SingleLinkedList) bool {
	head := l.Head
	if head == nil {
		return false
	}
	if head.Next == nil {
		return true
	}
	var slow, fast *SingleLinkedNode = head, head
	var prev, temp *SingleLinkedNode = nil, nil

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		temp = slow.Next
		slow.Next = prev
		prev = slow
		slow = temp
	} // 快的先跑完,同时反转了一半链表,剪短

	if fast != nil {
		slow = slow.Next // 处理余数，跨过中位数
		// prev 增加中 2 -> 1 -> nil
	}

	var l1 *SingleLinkedNode = prev
	var l2 *SingleLinkedNode = slow

	for l1 != nil && l2 != nil && l1.Val == l2.Val {
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}
