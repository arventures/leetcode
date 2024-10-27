package _133

// Node Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	cloned := make(map[*Node]*Node)

	var dfs func(*Node) *Node

	dfs = func(n *Node) *Node {
		if cloneNode, found := cloned[n]; found {
			return cloneNode
		}

		copy := &Node{Val: n.Val}
		cloned[n] = copy

		for _, neighbor := range n.Neighbors {
			copy.Neighbors = append(copy.Neighbors, dfs(neighbor))
		}

		return copy
	}

	return dfs(node)
}
