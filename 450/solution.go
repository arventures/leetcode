package _450

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func DeleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Right == nil {
			return root.Left
		}
		if root.Left == nil {
			return root.Right
		}

		cur := root.Right

		for cur.Left != nil {
			cur = cur.Left
		}

		root.Val = cur.Val

		root.Right = deleteNode(root.Right, cur.Val)

	}

	return root
}
