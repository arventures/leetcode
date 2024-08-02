package main

import (
	"container/heap"
)

// KthLargest struct using a min-heap
type KthLargest struct {
	k    int
	heap MinHeap
}

// MinHeap is a custom min-heap
type MinHeap []int

// Implement heap.Interface for MinHeap
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Constructor initializes the KthLargest object with k and the initial numbers.
func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{k: k}
	heap.Init(&kl.heap)
	for _, num := range nums {
		kl.Add(num)
	}
	return kl
}

// Add method adds a new number to the stream and returns the kth largest element.
func (kl *KthLargest) Add(val int) int {
	heap.Push(&kl.heap, val)
	if kl.heap.Len() > kl.k {
		heap.Pop(&kl.heap)
	}
	return kl.heap[0]
}
