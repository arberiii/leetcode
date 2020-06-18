package replaceElements

func replaceElements(arr []int) []int {
    max := arr[len(arr)-1]
    arr[len(arr)-1] = -1
    
    if len(arr) == 1 {
        return arr
    }
    
    for i := len(arr) - 2; i >= 0; i-- {
        if arr[i] > max {
            arr[i], max = max, arr[i] 
        } else {
            arr[i] = max
        }
    }
    
    return arr
}
