func lengthOfLongestSubstring(s string) int {
    m := make(map[rune]int)
    longestSubstring, currentSubstring, startOfSubstring := 0, 0, 0
    for index, char := range s {
        value, ok := m[char]
        if ok && value >= startOfSubstring {
            startOfSubstring = value + 1
            currentSubstring = index - startOfSubstring
        }
        currentSubstring += 1
        if longestSubstring < currentSubstring {
            longestSubstring = currentSubstring
        }
        m[char] = index
    }
    
    return longestSubstring
}
