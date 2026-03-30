func twoSum(nums []int, target int) []int {
    hash := make(map[int]int)
    for i, num := range nums {
        difference := target - num
        if idx, ok := hash[difference]; ok {
            return []int{idx, i}
        }
        hash[num] = i
    }
    return []int{}
}
