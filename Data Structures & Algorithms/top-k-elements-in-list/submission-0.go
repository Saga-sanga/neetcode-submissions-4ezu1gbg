import "slices"

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)

	for _, num := range nums {
		count[num]++
	}

	arr := make([][2]int, 0, len(count))
	for num, cnt := range count {
		arr = append(arr, [2]int{cnt, num})
	}

	slices.SortFunc(arr, func(a, b [2]int) int {
		return b[0] - a[0]
	})

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = arr[i][1]
	}

	return res
}
