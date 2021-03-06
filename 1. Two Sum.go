package twoSum

func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i := range nums {
        if v, ok := m[target-nums[i]]; ok {
            return []int{v, i}
        } else {
            m[nums[i]] = i
        }
    }
    return []int{-1, -1}
}
