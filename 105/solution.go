package _105

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	mid := index(inorder, root.Val)

	root.Left = BuildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = BuildTree(preorder[mid+1:], inorder[mid+1:])

	return root
}

func index(arr []int, val int) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}
	return -1
}
