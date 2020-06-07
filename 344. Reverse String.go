package reverseString

func reverseString(s []byte)  {
    if len(s) == 0 {
		return
	}
	length := len(s) - 1
	for i := 0; i <= length/2; i++ {
		s[i], s[length-i] = s[length-i], s[i]
	}
}
