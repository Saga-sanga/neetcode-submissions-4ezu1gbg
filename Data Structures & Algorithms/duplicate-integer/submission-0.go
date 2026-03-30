func hasDuplicate(nums []int) bool {
    hash := make(map[int]bool)
    for _, num := range nums {
        hash[num] = true
    }
    if len(hash) == len(nums) {
        return false
    }
    return true
}
