package getHint

import "strconv"

func getHint(secret string, guess string) string {
    bulls := 0
    cows := 0
    used := make([]bool, len(secret))
    for i := range secret {
        if secret[i] == guess[i] {
            bulls++
            used[i] = true
        }
    }
    
    m := make(map[rune]int)
    for i, ch := range secret {
        if used[i] {
            continue
        }
        
        if _, ok := m[ch]; !ok {
            m[ch] = 1
        } else {
            m[ch]++
        }
    }
    
    for i, ch := range guess {
        if used[i] {
            continue
        }
        
        if v, ok := m[ch]; ok {
            if v > 0 {
                cows++
                m[ch]--
            }
        }
    }
    
    return strconv.Itoa(bulls)+"A"+strconv.Itoa(cows)+"B"
}
