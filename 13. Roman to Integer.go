func romanToInt(s string) int {
    m := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    
    total := 0
    for i := range s {
        if len(s) == i + 1 {
            total += m[s[i]]
            return total
        }

        if m[s[i]] < m[s[i+1]] {
            total -= m[s[i]]
        } else {
            total += m[s[i]]
        }
    }
    
    return total
}
