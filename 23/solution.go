package _23

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func MergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	for len(lists) > 1 {
		mergedLists := []*ListNode{}
		for i := 0; i < len(lists); i += 2 {
			if i+1 < len(lists) {
				mergedLists = append(mergedLists, mergeTwoLists(lists[i], lists[i+1]))
			} else {
				mergedLists = append(mergedLists, lists[i])
			}
		}
		lists = mergedLists
	}

	return lists[0]
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	output := &ListNode{}
	current := output

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next // Move the current pointer forward
	}

	if l1 != nil {
		current.Next = l1
	}
	if l2 != nil {
		current.Next = l2
	}

	return output.Next
}
