package longestCommonPrefix

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    currentBest := len(strs[0])
    for _, word := range strs {
        for i := 0; i < currentBest; i++ {
            if i == len(word) {
                currentBest = i
                break
            } else if i == currentBest {
                break
            } else if word[i] != strs[0][i] {
                currentBest = i
                break
            }
        }
    }
    return strs[0][:currentBest]
}
