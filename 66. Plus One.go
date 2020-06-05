package plusOne

func plusOne(digits []int) []int {
    length := len(digits) + 1
    for i := range digits {
        if digits[i] != 9 {
            length = len(digits)
        }
    }
    
    if length == len(digits) + 1 {
        return plusOneNines(len(digits))
    }
    
    remainder := true
    for i := len(digits) - 1; i >= 0; i-- {
        if remainder {
            if digits[i] + 1 == 10 {
                digits[i] = 0
                remainder = true
            } else {
                digits[i] = digits[i] + 1
                remainder = false
            }
        } else {
            digits[i] = digits[i]
            remainder = false
        }
    } 
    
    return digits
}

func plusOneNines(length int) []int{
    result := make([]int, length+1)
    result[0] = 1
    return result
}
