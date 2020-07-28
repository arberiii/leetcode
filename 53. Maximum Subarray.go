package maxSubArray

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	ret := nums[0]
	sum := 0
	for i := 0; i < len(nums); i++ {
		val := sum + nums[i]
		ret = max(ret, val)
		if val <= 0 {
			sum = 0
		} else {
			sum = val
		}
	}

	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
