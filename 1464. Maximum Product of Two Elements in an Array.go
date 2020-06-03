package maxProduct

func maxProduct(nums []int) int {
    var max1, max2 int
    if nums[0] > nums[1] {
        max1, max2 = nums[0], nums[1]
    } else {
        max2, max1 = nums[0], nums[1]
    }
    
    for _, num := range nums[2:] {
        if num > max2 {
            if num > max1 {
                max2 = max1
                max1 = num
            } else {
                max2 = num
            }
        }
    }
    
    return (max1-1)*(max2-1)
}
