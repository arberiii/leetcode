package generate

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	done := make(chan bool)
    
    if numRows == 0 {
        return result
    }

	result[0] = []int{1}
	if numRows == 1 {
		return result
	}

	result[1] = []int{1, 1}

	for i := range result[2:] {
		go func(i int){
			result[i+2] = make([]int, i+3)
			generateRow(i+2, result[i+2])
			done <- true
		}(i)
	}

	for _ = range result[2:] {
		<-done
	}
	return result
}

func generateRow(mul int, row []int) {
	row[0] = 1
	div := 1
	for i := range row[1:] {
		row[i+1] = row[i] * mul
		row[i+1] /= div
		mul--
		div++
	}
}
