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

type KthLargest struct {
	k    int
	heap *MinHeap
}

func Constructor(k int, nums []int) KthLargest {
	h := MinHeap(nums)
	heap.Init(&h)
	kl := KthLargest{k: k, heap: &h}
	for kl.heap.Len() > k {
		heap.Pop(kl.heap)
	}
	return kl
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.heap, val)
	if this.heap.Len() > this.k {
		heap.Pop(this.heap)
	}
	return (*this.heap)[0]
}

// type KthLargest struct {
// 	k       int
// 	minHeap []int
// }

// func Constructor(k int, nums []int) KthLargest {
// 	for i := len(nums)/2 - 1; i >= 0; i-- {
// 		bubbleDown(i, nums)
// 	}
// 	for len(nums) > k {
// 		nums = pop(nums)
// 	}
// 	return KthLargest{
// 		k:       k,
// 		minHeap: nums,
// 	}
// }

// func pop(nums []int) []int {
// 	nums[0], nums[len(nums)-1] = nums[len(nums)-1], nums[0]
// 	nums = nums[:len(nums)-1]
// 	bubbleDown(0, nums)
// 	return nums
// }

// func bubbleDown(i int, nums []int) {
// 	for {
// 		smallest := i
// 		left := 2*i + 1
// 		right := 2*i + 2

// 		if left < len(nums) && nums[left] < nums[smallest] {
// 			smallest = left
// 		}
// 		if right < len(nums) && nums[right] < nums[smallest] {
// 			smallest = right
// 		}
// 		//smallest = 6
// 		if smallest == i {
// 			break
// 		}

// 		nums[i], nums[smallest] = nums[smallest], nums[i]
// 		i = smallest
// 	}
// }

// func bubbleUp(i int, nums []int) {
// 	for i > 0 {
// 		parent := (i - 1) / 2
// 		if nums[parent] <= nums[i] {
// 			break
// 		}
// 		nums[i], nums[parent] = nums[parent], nums[i]
// 		i = parent
// 	}

// }

// func (this *KthLargest) Add(val int) int {
// 	this.minHeap = append(this.minHeap, val)
// 	bubbleUp(len(this.minHeap)-1, this.minHeap)
// 	for len(this.minHeap) > this.k {
// 		this.minHeap = pop(this.minHeap)
// 	}
// 	return this.minHeap[0]
// }

// // 8 5 4 3 2
// // func (this *KthLargest) Add(val int) int {
// // }

// func main() {
// 	f := Constructor(3, []int{4, 5, 8, 2})
// 	fmt.Print(f.minHeap)
// }

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
