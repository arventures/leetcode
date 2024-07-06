package _199

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var response []int

	var queue []*TreeNode

	queue = append(queue, root)

	for len(queue) > 0 {
		length := len(queue)

		var lastNode *TreeNode

		for i := 0; i < length; i++ {

			node := queue[0]
			queue = queue[1:]

			lastNode = node

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)

			}
		}

		if lastNode != nil {
			response = append(response, lastNode.Val)
		}

	}

	return response
}
