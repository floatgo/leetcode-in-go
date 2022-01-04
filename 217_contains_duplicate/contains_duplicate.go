func containsDuplicate(nums []int) bool {
    m := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        if _, ok := m[nums[i]]; ok {
            m[nums[i]] += 1
        } else {
            m[nums[i]] = 1
        }
    }
    
    for _, v := range m {
        if v > 1 {
            return true
        }
    }
    return false
}
