package getRow

func getRow(rowIndex int) []int {
    if rowIndex == 0 {
        return []int{1}
    }
    result := make([]int, rowIndex+1)
    result[0] = 1
	div := 1
	for i := range result[1:] {
		result[i+1] = result[i] * rowIndex
		result[i+1] /= div
		rowIndex--
		div++
	}
    return result
}
