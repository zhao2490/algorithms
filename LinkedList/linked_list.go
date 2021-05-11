package LinkedList

/*
单链表反转 时间复杂度 O(n)
*/
func ReverseLinkedList(node *SingleLinkedNode) {
	var pre *SingleLinkedNode = nil
	cur := node
	for cur.Next != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	node = pre
}

/*
链表中环的检测
*/
func HasCycle(node *SingleLinkedNode) bool {
	if node != nil {
		slow, fast := node, node
		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
			if slow == fast {
				return true
			}
		}
	}
	return false
}

/*
两个有序链表的合并
*/
func MergeSortedLinkedList(l1 *SingleLinkedNode, l2 *SingleLinkedNode) *SingleLinkedNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	pre := &SingleLinkedNode{}
	cur1 := l1
	cur2 := l2
	for cur1 != nil && cur2 != nil {
		if cur1.Val < cur2.Val {
			pre.Next = cur1
			cur1 = cur1.Next
		} else {
			pre.Next = cur2
			cur2 = cur2.Next
		}
		pre = pre.Next
	}
	if cur1 != nil {
		pre.Next = cur1
	}
	if cur2 != nil {
		pre.Next = cur2
	}
	return pre.Next
}

/*
删除链表倒数第n个节点
*/
func DeleteBottomN(node *SingleLinkedNode, n int) {
	if n <= 0 || node == nil {
		return
	}
	fast := node
	for i := 1; i <= n && fast != nil; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return
	}

	slow := node
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
}

/*
求链表的中间节点
*/
func FindMiddleNode(node *SingleLinkedNode) *SingleLinkedNode {
	if node == nil {
		return nil
	}
	var slow, fast *SingleLinkedNode = node, node
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
