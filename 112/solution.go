package _112

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return checkPathSum(root, 0, targetSum)
}

func checkPathSum(node *TreeNode, currentSum int, targetSum int) bool {
	if node == nil {
		return false
	}

	currentSum += node.Val

	if node.Left == nil && node.Right == nil {
		return currentSum == targetSum
	}

	leftHasPathSum := checkPathSum(node.Left, currentSum, targetSum)
	rightHasPathSum := checkPathSum(node.Right, currentSum, targetSum)

	return leftHasPathSum || rightHasPathSum
}
