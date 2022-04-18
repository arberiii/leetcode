package generateMatrix

func generateMatrix(n int) [][]int {
    a := make([][]int, n)
    for i := range a {
        a[i] = make([]int, n)
    }
    
    a = writeSubmatrix(a, 0, 0, n, 0)
    
    
    return a
}

func writeSubmatrix(matrix [][]int, row int, column int, length int, startValue int) [][]int {
    for i := 0; i < length; i++ {
        matrix[row][i + column] = i + 1 + startValue
    }
    
    for i := 1; i < length; i++ {
        matrix[row + i][length - 1 + column] = length + i + startValue
    }
    
    for i := 1; i < length; i++ {
        matrix[length - 1 + row][length - 1 - i + column] = 2 * length + i - 1 + startValue
    }
    
    for i := 1; i < length - 1; i++ {
        matrix[length - i - 1 + row][column] = 3 * length + i - 2 + startValue
    }
    
    if length <= 2 {
        return matrix
    }
    matrix = writeSubmatrix(matrix, row + 1, column + 1, length - 2, 4 * length - 4 + startValue)
    return matrix
}
