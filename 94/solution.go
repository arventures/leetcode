package _94

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode) []int {
	var response []int
	return traverse(root, response)
}

func traverse(node *TreeNode, response []int) []int {
	if node == nil {
		return response
	}

	response = traverse(node.Left, response)
	response = append(response, node.Val)
	response = traverse(node.Right, response)

	return response
}
