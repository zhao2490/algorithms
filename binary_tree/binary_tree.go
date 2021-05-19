package binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

// 前序遍历
func preOrderTraversal(node *TreeNode, res []int) []int {
	if node == nil {
		return res
	}
	res = append(res, node.Val)
	res = preOrderTraversal(node.Left, res)
	res = preOrderTraversal(node.Right, res)
	return res
}

// 使用栈实现前序遍历
func preOrderTraversal2(node *TreeNode) []int {
	var stack []*TreeNode
	var res []int
	stack = append(stack, node)
	for len(stack) != 0 {
		e := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, e.Val)
		if e.Right != nil {
			stack = append(stack, e.Right)
		}
		if e.Left != nil {
			stack = append(stack, e.Left)
		}
	}
	return res
}

// 使用递归实现中序遍历
func inOrderTraversal(node *TreeNode, res []int) []int {
	if node == nil {
		return res
	}
	res = inOrderTraversal(node.Left, res)
	res = append(res, node.Val)
	res = inOrderTraversal(node.Right, res)
	return res
}

// 使用栈实现中序遍历
func inOrderTraversal2(node *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		if len(stack) > 0 {
			p := stack[len(stack)-1]
			res = append(res, p.Val)
			node = p.Right
			stack = stack[:len(stack)-1]
		}
	}
	return res
}

func postOrderTraversal(node *TreeNode, res []int) []int {
	if node == nil {
		return res
	}
	res = postOrderTraversal(node.Left, res)
	res = postOrderTraversal(node.Right, res)
	res = append(res, node.Val)
	return res
}

// 使用栈实现后序遍历
func postOrderTraversal2(node *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var prev *TreeNode
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		if len(stack) > 0 {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if node.Right == nil || node.Right == prev {
				res = append(res, node.Val)
				prev = node
				node = nil
			} else {
				stack = append(stack, node)
				node = node.Right
			}
		}
	}
	return res
}
