package isMonotonic

func isMonotonic(A []int) bool {
    if len(A) == 1 {
        return true
    }
    
    arrPointer := 1
    seqIncreasing := true
    for {
        if A[arrPointer] > A[0] {
            break
        } else if A[arrPointer] < A[0] {
            seqIncreasing = false
            break
        }
        
        arrPointer++
        if len(A) == arrPointer {
            return true
        }
    }
    
    for i := range A[:len(A)-1] {
        if seqIncreasing {
            if A[i] > A[i+1] {
                return false
            }
        } else {
            if A[i] < A[i+1] {
                return false
            }
        }
    }
    return true
}
