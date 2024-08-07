package _1046

import "container/heap"

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	oldHeap := *h
	n := len(oldHeap)
	x := oldHeap[n-1]
	*h = oldHeap[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	h := &MaxHeap{}
	heap.Init(h)

	for _, stone := range stones {
		heap.Push(h, stone)
	}

	for h.Len() > 1 {
		first := heap.Pop(h).(int)
		second := heap.Pop(h).(int)

		if first != second {
			heap.Push(h, first-second)
		}
	}

	if h.Len() == 1 {
		return heap.Pop(h).(int)
	}

	return 0
}
