func merge(intervals [][]int) [][]int {
    var ret [][]int
    intervals = sortIntervals(intervals)
    if len(intervals) == 0 {
        return intervals
    }
    var currentInterval = intervals[0]
    for i := 1; i < len(intervals); i++ {
        if (currentInterval[1] < intervals[i][0]) {
            ret = append(ret, currentInterval)
            currentInterval = intervals[i]
        } else {
            currentInterval = mergeTwoIntervals(currentInterval, intervals[i])
        }
    }
    ret = append(ret, currentInterval)
    return ret
}

func mergeTwoIntervals(first, second []int) []int {
    if (first[1] < second[1]){
        return []int{first[0], second[1]}
    }
    return []int{first[0], first[1]}
}

func sortIntervals(intervals [][]int) [][]int {
    sort.SliceStable(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    return intervals
}
