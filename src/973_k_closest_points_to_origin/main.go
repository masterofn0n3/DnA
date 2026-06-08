package main

import "container/heap"

type MinHeap [][]int

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	return h[i][0]*h[i][0]+h[i][1]*h[i][1] < h[j][0]*h[j][0]+h[j][1]*h[j][1]
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)   { *h = append(*h, x.([]int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x

}

func kClosest(points [][]int, k int) [][]int {
	result := [][]int{}
	h := MinHeap(points)
	heap.Init(&h)

	for i := 0; i < k; i++ {

		result = append(result, heap.Pop(&h).([]int))
	}

	return result
}
