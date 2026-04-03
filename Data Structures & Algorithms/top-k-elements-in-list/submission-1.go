type Item struct {
	value int
	count int
}

type MinHeap []Item

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].count < h[j].count
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}

	h := &MinHeap{}
	heap.Init(h)

	for num, count := range counts {
		heap.Push(h, Item{value: num, count: count})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]int, k)
	for i := k-1; i >= 0; i-- {
		item := heap.Pop(h).(Item)
		res[i] = item.value
	}

	return res
}
