package addDigits

func addDigits(num int) int {
    if (num == 0) {
	    return 0;
	} else if (num % 9 == 0) {
	    return 9;
	}
	return num % 9;
}
