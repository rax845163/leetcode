package q202

func isHappy(n int) bool {
	sum := 0
	for n > 0 {
		r := n % 10
		sum += r * r
		n /= 10
	}
	return false
}
