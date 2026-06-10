package main

import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := MinHeap([]int{})
	heap.Init(&h)
	i := 0
	for i < k {
		heap.Push(&h, nums[i])
		i++
	}
	for i < len(nums) {
		heap.Push(&h, nums[i])
		heap.Pop(&h)
		i++
	}

	return h[0]
}
