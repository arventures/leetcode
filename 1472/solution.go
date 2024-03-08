package _1472

type Node struct {
	val  string
	prev *Node
	next *Node
}

type BrowserHistory struct {
	curr *Node
}

func Constructor(homepage string) BrowserHistory {
	curr := &Node{
		val: homepage,
	}

	return BrowserHistory{curr: curr}
}

func (this *BrowserHistory) Visit(url string) {
	newNode := &Node{val: url, prev: this.curr}
	this.curr.next = newNode
	this.curr = newNode
}

func (this *BrowserHistory) Back(steps int) string {
	for this.curr.prev != nil && steps > 0 {
		this.curr = this.curr.prev
		steps--
	}
	return this.curr.val
}

func (this *BrowserHistory) Forward(steps int) string {
	for this.curr.next != nil && steps > 0 {
		this.curr = this.curr.next
		steps--
	}
	return this.curr.val
}
