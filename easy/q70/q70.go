package q70

var cache = map[int]int{}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	var a, b int
	var exist bool
	if a, exist = cache[n-1]; !exist {
		a = climbStairs(n - 1)
		cache[n-1] = a
	}
	if b, exist = cache[n-2]; !exist {
		b = climbStairs(n - 2)
		cache[n-2] = b
	}

	return a + b
}
