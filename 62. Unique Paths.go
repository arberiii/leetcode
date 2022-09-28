func uniquePaths(m int, n int) int {
    return Choose(m+n-2, n-1)
}

func Choose(n, k int) int {
	if k > n {
		panic("Choose: k > n")
	}
	if k < 0 {
		panic("Choose: k < 0")
	}
	if n <= 1 || k == 0 || n == k {
		return 1
	}
	if newK := n - k; newK < k {
		k = newK
	}
	if k == 1 {
		return n
	}
	ret := int(n - k + 1)
	for i, j := ret+1, int(2); j <= k; i, j = i+1, j+1 {
		ret = ret * i / j
	}
	return ret
}
