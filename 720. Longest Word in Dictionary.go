package longestWord

func longestWord(words []string) string {
    m := make(map[string]bool)
    for _, word := range words {
        m[word] = true
    }
    longestWord := ""
    
    for _, word := range words {
        canBeBuild := true
        for i := 1; i < len(word); i++ {
            if !canBeBuild {
                continue
            }
            
            ok, _ := m[word[:i]]
            if !ok {
                canBeBuild = false
            }
        }
        
        if canBeBuild && len(word) >= len(longestWord) {
            if len(word) > len(longestWord) {
                longestWord = word
            } else {
                if len(word) == len(longestWord) && word < longestWord {
                    longestWord = word
                }
            }
        }
    }
    
    return longestWord
}
