package q66

func plusOne(digits []int) []int {
	digits[len(digits)-1] = digits[len(digits)-1] + 1
	c := 0
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = digits[i] + c
		c = digits[i] / 10
		digits[i] = digits[i] % 10
	}
	if c > 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
