func twoSum(nums []int, target int) []int {
    hash := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        difference := target - nums[i]
        if idx, ok := hash[difference]; ok {
            return []int{idx, i}
        }
        hash[nums[i]] = i
    }
    return []int{0,0}
}
