package main

import "container/heap"

type MaxHeap []int

func (m MaxHeap) Len() int            { return len(m) }
func (m MaxHeap) Less(i, j int) bool  { return m[i] > m[j] }
func (m MaxHeap) Swap(i, j int)       { m[i], m[j] = m[j], m[i] }
func (m *MaxHeap) Push(x interface{}) { *m = append(*m, x.(int)) }
func (m *MaxHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]

	return x

}

func lastStoneWeight(stones []int) int {
	h := MaxHeap(stones)
	heap.Init(&h)

	for len(h) > 1 {
		first := heap.Pop(&h).(int)
		second := heap.Pop(&h).(int)

		if first > second {
			heap.Push(&h, first-second)
		}
	}

	return h[0]
}
