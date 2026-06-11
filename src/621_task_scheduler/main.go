package main

// func leastInterval(tasks []byte, n int) int {
// 	taskMap := make(map[byte]int)
// 	maxFreq := 0
// 	countMaxFreq := 0
// 	for _, v := range tasks {
// 		taskMap[v]++
// 		maxFreq = max(maxFreq, taskMap[v])
// 	}
// 	for _, v := range taskMap {
// 		if v == maxFreq {
// 			countMaxFreq++
// 		}
// 	}

// 	return max((maxFreq-1)*(n+1)+countMaxFreq, len(tasks))

// }
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

// func leastInterval(tasks []byte, n int) int {
// 	taskMap := make(map[byte]int)
// 	maxFreq := 0
// 	for _, v := range tasks {
// 		taskMap[v]++
// 		maxFreq = max(maxFreq, taskMap[v])
// 	}

// 	h := MaxHeap([]int{})
// 	heap.Init(&h)

// 	for _, v := range taskMap {
// 		heap.Push(&h, v)
// 	}

// 	taskQueue := make([]struct {
// 		int
// 		byte
// 	}, n)

// 	for len(h) > 0 {
// 		currentTaskFreq := heap.Pop(&h)

// 	}

// }
