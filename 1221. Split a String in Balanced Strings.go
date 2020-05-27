package balancedStringSplit

func balancedStringSplit(s string) int {
    currentLetter := s[0]
    result := 0
    counter := 0
    for i := range s {
        if counter == 0 {
            currentLetter = s[i]
        }
        
        if s[i] == currentLetter {
            counter = counter + 1
        } else {
            counter = counter - 1
            if counter == 0 {
                result = result + 1
            }
        }
    }
    
    return result
}
