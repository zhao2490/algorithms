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
func perOrderTraversal2(node *TreeNode) []int {
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

func inOrderTraversal(node *TreeNode, res []int) []int {
	if node == nil {
		return res
	}
	res = inOrderTraversal(node.Left, res)
	res = append(res, node.Val)
	res = inOrderTraversal(node.Right, res)
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
