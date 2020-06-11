package peakIndexInMountainArray

func peakIndexInMountainArray(A []int) int {
    peak := 0
    for i := range A {
        if A[i] < A[peak] {
            return peak
        } else {
            peak = i
        }
    }
    
    return peak
}
