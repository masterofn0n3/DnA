package main

import (
	"container/heap"
)

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

type MedianFinder struct {
	maxHeap MaxHeap
	minHeap MinHeap
}

func Constructor() MedianFinder {
	maxHeap := MaxHeap([]int{})
	minHeap := MinHeap([]int{})
	heap.Init(&maxHeap)
	heap.Init(&minHeap)
	return MedianFinder{
		maxHeap,
		minHeap,
	}
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.maxHeap) == 0 && len(this.minHeap) == 0 {
		heap.Push(&this.maxHeap, num)
		return
	}
	if num <= this.maxHeap[0] {
		heap.Push(&this.maxHeap, num)
	} else {
		heap.Push(&this.minHeap, num)
	}

	if len(this.maxHeap)-len(this.minHeap) > 1 {
		heap.Push(&this.minHeap, heap.Pop(&this.maxHeap))
	} else if len(this.minHeap)-len(this.maxHeap) > 1 {
		heap.Push(&this.maxHeap, heap.Pop(&this.minHeap))
	}

}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.maxHeap) > len(this.minHeap) {
		return float64(this.maxHeap[0])
	} else if len(this.maxHeap) == len(this.minHeap) {
		return (float64(this.maxHeap[0]) + float64(this.minHeap[0])) / 2
	} else {
		return float64(this.minHeap[0])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
