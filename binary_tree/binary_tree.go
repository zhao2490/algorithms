package binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrderTraversal(node *TreeNode, res []int) []int {
	if node == nil {
		return res
	}
	res = append(res, node.Val)
	res = preOrderTraversal(node.Left, res)
	res = preOrderTraversal(node.Right, res)
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
