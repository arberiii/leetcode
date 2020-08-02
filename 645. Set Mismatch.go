package findErrorNums

func findErrorNums(nums []int) []int {
    count := make([]int, len(nums))
    
    for i := range nums {
        count[nums[i]-1] += 1
    }
    
    n1 := 0
    n2 := 0
    for i := range count {
        if count[i] == 2 {
            n1 = i + 1
        } else if count[i] == 0 {
            n2 = i + 1
        }
    }
    
    return []int{n1, n2}
}
