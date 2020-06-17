package runningSum

func runningSum(nums []int) []int {
    for i := range nums[1:] {
        nums[i+1] = nums[i+1] + nums[i]
    }
    return nums
}
