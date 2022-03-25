package isPalindrome

func isPalindrome(x int) bool {
    if (x < 0) {
        return false
    }
    
    digits := 1
    xTemp := x
    for !(xTemp < 10) {
        xTemp = xTemp / 10
        digits += 1
    }
    
    y := intPow(10, digits - 1)
    for i := 1; i <= digits / 2; i++ {
        firstDigit := x / y
        lastDigit := x % 10
        
        if firstDigit != lastDigit {
            return false
        }
        
        x = x % y
        x = x / 10
        y = y / 100
    }
    
    return true
}

func intPow(n, m int) int {
    if m == 0 {
        return 1
    }
    result := n
    for i := 2; i <= m; i++ {
        result *= n
    }
    return result
}
