package solution

import (
	"container/heap"
	"strconv"
)

type Heap []int

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Less(i int, j int) bool {
	return strconv.Itoa((*h)[i]) > strconv.Itoa((*h)[j])
}

func (h *Heap) Swap(i int, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() any {
	ret := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return ret
}

func findKthNumber(n int, k int) int {
	h := make(Heap, 0)
	for i := 1; i <= n; i++ {
		heap.Push(&h, i)
		if len(h) > k {
			heap.Pop(&h)
		}
	}
	return heap.Pop(&h).(int)
}
