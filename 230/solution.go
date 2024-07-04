package _230

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func KthSmallest(root *TreeNode, k int) int {

	var response []int

	response = traverse(root, response)

	return response[k-1]

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
