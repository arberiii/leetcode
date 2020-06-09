package sortArrayByParity

func sortArrayByParity(A []int) []int {
    p := 0
    pOdd := len(A) - 1
    for {
        if A[p] % 2 == 0 {
            p++
        } else {
            A[p], A[pOdd] = A[pOdd], A[p]
            pOdd--
        }
        
        if p >= pOdd {
            return A
        }
    }
}
