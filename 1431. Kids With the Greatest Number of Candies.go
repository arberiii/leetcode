package kidsWithCandies

func kidsWithCandies(candies []int, extraCandies int) []bool {
    result := make([]bool, len(candies))
    m := max(candies)
    for i := range result {
        if candies[i] + extraCandies >= m {
            result[i] = true
        } else {
            result[i] = false
        }
    }
    return result
}

func max(candies []int) int {
    result := candies[0]
    for i := range candies {
        if candies[i] > result {
            result = candies[i]
        }
    }
    return result
}
