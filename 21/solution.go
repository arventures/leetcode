package _21

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next // Move forward in list1.
		} else {
			tail.Next = list2
			list2 = list2.Next // Move forward in list2.
		}
		tail = tail.Next // Move forward in the merged list.
	}

	// Attach any remaining elements from either list to the merged list.
	if list1 != nil {
		tail.Next = list1
	} else if list2 != nil {
		tail.Next = list2
	}

	return dummy.Next // Return the start of the merged list, skipping the dummy.
}
