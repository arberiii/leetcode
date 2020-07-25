package trailingZeroes

func trailingZeroes(n int) int {
    sum := 0
    div := 5
    for {
        val := n / div
        if val == 0 {
            break
        }
        sum += val
        div *= 5
    }
    return sum
}
