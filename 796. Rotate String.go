func rotateString(A string, B string) bool {
    if len(A) != len(B) {
        return false
    }
    
    for _ = range A {
        if A == B {
            return true
        }

        A = shiftString(A)
    }
    
    return A == B
}

func shiftString(A string) string {
    return A[1:] + string(A[0])
}
