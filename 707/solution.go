package _707

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

type MyLinkedList struct {
	Head *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	curr := this.Head

	for i := 0; i < index && curr != nil; i++ {
		curr = curr.Next
	}
	if curr == nil {
		return -1
	}
	return curr.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{Val: val}

	if this.Head != nil {
		node.Next = this.Head
		this.Head.Prev = node
	}

	this.Head = node

}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{
		Val:  val,
		Next: nil,
	}

	if this.Head == nil {
		this.Head = node
		return
	}

	curr := this.Head

	for curr.Next != nil {
		curr = curr.Next
	}

	node.Prev = curr
	curr.Next = node
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.AddAtHead(val)
		return
	}
	curr := this.Head
	for i := 0; i < index-1 && curr != nil; i++ {
		curr = curr.Next
	}
	if curr == nil {
		return
	}
	if curr.Next == nil {
		this.AddAtTail(val)
		return
	}
	node := &Node{Val: val}
	nextNode := curr.Next
	curr.Next = node
	node.Prev = curr
	node.Next = nextNode
	nextNode.Prev = node
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || this.Head == nil {
		return
	}
	curr := this.Head
	if index == 0 {
		this.Head = curr.Next
		if this.Head != nil {
			this.Head.Prev = nil
		}
		return
	}
	for i := 0; i < index && curr != nil; i++ {
		curr = curr.Next
	}
	if curr == nil {
		return
	}
	if curr.Prev != nil {
		curr.Prev.Next = curr.Next
	}
	if curr.Next != nil {
		curr.Next.Prev = curr.Prev
	}
}
