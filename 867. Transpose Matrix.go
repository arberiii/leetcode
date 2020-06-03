package transpose

func transpose(A [][]int) [][]int {
    // init
    AT := make([][]int, len(A[0]))
    for i := range AT {
        AT[i] = make([]int, len(A))
    }
    
    for i := range A {
        for j := range A[i] {
            AT[j][i] = A[i][j] 
        }
    }
    
    return AT
}
