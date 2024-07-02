package _450

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DeleteNode(root *TreeNode, key int) *TreeNode {

	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = DeleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = DeleteNode(root.Right, key)
	} else {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			//Case where node has two nodes
			minNode := MinNode(root.Right)
			root.Val = minNode.Val
			root.Right = DeleteNode(root.Right, minNode.Val)

		}
	}
	return root
}

func MinNode(root *TreeNode) *TreeNode {
	if root.Left == nil {
		return root
	} else {
		return MinNode(root.Left)
	}
}
