package _102

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	var response [][]int

	var queue []*TreeNode

	queue = append(queue, root)

	for len(queue) > 0 {
		length := len(queue)
		var level []int

		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		response = append(response, level)

	}

	return response
}
