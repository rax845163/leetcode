package q645

func findErrorNums(nums []int) []int {
	xor := 0
	for i, num := range nums {
		xor = (i + 1) ^ num ^ xor
	}

	a := 1
	for xor != 0 {
		b := a ^ xor
		if a^b == xor {
			return []int{a, b}
		}
		a++
	}
	return []int{}
}
