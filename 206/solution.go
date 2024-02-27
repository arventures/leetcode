package _206

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {

	var prev *ListNode = nil

	for head != nil {

		head.Next, prev, head = prev, head, head.Next

	}

	return prev

}
