package binary_tree

import "fmt"

type Node struct {
	Data  interface{}
	Left  *Node
	Right *Node
}

func NewNode(data interface{}) *Node {
	return &Node{Data: data}
}

/*
二叉查找树
*/
type BinarySearchTree struct {
	root *Node
	// 对比函数 0：v==nodeV，正数：v>nodeV，负数：v<nodeV
	compareFunc func(v, nodeV interface{}) int
}

func NewBinarySearchTree(rootV interface{}, compareFunc func(v, nodeV interface{}) int) *BinarySearchTree {
	if compareFunc == nil {
		return nil
	}
	return &BinarySearchTree{root: NewNode(rootV), compareFunc: compareFunc}
}

func (t *BinarySearchTree) Find(v interface{}) *Node {
	p := t.root
	for p != nil {
		compareResult := t.compareFunc(v, p.Data)
		if compareResult == 0 {
			return p
		} else if compareResult > 0 { // v > nodeV
			p = p.Right
		} else { // v < nodeV
			p = p.Left
		}
	}
	return nil
}

func (t *BinarySearchTree) Insert(v interface{}) bool {
	p := t.root
	for p != nil {
		compareResult := t.compareFunc(v, p.Data)
		if compareResult == 0 {
			return false
		} else if compareResult > 0 {
			if p.Right == nil {
				p.Right = NewNode(v)
				break
			}
			p = p.Right
		} else {
			if p.Left == nil {
				p.Left = NewNode(v)
				break
			}
			p = p.Left
		}
	}
	return true
}

func (t *BinarySearchTree) Delete(v interface{}) bool {
	var pp *Node = nil // 记录需要删除节点的父节点
	p := t.root
	deleteLeft := false
	for p != nil {
		compareResult := t.compareFunc(v, p.Data)
		if compareResult > 0 {
			pp = p
			p = p.Right
			deleteLeft = false
		} else if compareResult < 0 {
			pp = p
			p = p.Left
			deleteLeft = true
		} else {
			break
		}
	}

	if p == nil { // 需要删除的节点不存在
		return false
	} else if p.Left == nil && p.Right == nil { // 删除的是一个叶子节点
		if pp != nil {
			if deleteLeft {
				pp.Left = nil
			} else {
				pp.Right = nil
			}
		} else { // 根结点需要被删除
			t.root = nil
		}
	} else if p.Right != nil { // 删除的是一个有右孩子，不一定有左孩子的节点
		// 找到p节点右孩子的最小节点
		pq := p
		q := p.Right // 向右走一步
		fromRight := true
		for q.Left != nil { // 向左走到底
			pq = q
			q = q.Left
			fromRight = false
		}
		if fromRight {
			pq.Right = nil
		} else {
			pq.Left = nil
		}
		q.Left = p.Left
		q.Right = p.Right
		if nil == pp { //根节点被删除
			t.root = q
		} else {
			if deleteLeft {
				pq.Left = q
			} else {
				pq.Right = q
			}
		}
	} else { //删除的是一个只有左孩子的节点
		if nil != pp {
			if deleteLeft {
				pp.Left = p.Left
			} else {
				pp.Right = p.Left
			}
		} else {
			if deleteLeft {
				t.root = p.Left
			} else {
				t.root = p.Left
			}
		}
	}

	return true

}

func (t *BinarySearchTree) Min() *Node {
	p := t.root
	for nil != p.Left {
		p = p.Left
	}
	return p
}

func (t *BinarySearchTree) Max() *Node {
	p := t.root
	for nil != p.Right {
		p = p.Right
	}
	return p
}

func (t *BinarySearchTree) InOrderTraversal() {
	var stack []*Node
	var res []interface{}
	p := t.root
	for p != nil || len(stack) > 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		if len(stack) > 0 {
			p = stack[len(stack)-1]
			res = append(res, p.Data)
			stack = stack[:len(stack)-1]
			p = p.Right
		}
	}
	fmt.Println(res)
}
