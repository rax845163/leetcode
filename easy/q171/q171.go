package q171

func titleToNumber(s string) int {
	e := 1
	sum := 0
	for i := len(s) - 1; i >= 0; i-- {
		c := int(s[i]-'A') + 1
		sum += (26 * e) * c
		e *= 26
	}
	return sum
}

// func titleToNumber(s string) int {
// 	e := len(s) - 1
// 	sum := 0
// 	for i := 0; i < len(s); i++ {
// 		c := int(s[i]-'A') + 1
// 		sum += pow(26, e-i) * c
// 	}
// 	return sum
// }

// func pow(base, e int) int {
// 	if e == 0 {
// 		return 1
// 	}
// 	s := 1
// 	for e > 0 {
// 		s *= base
// 		e--
// 	}
// 	return s
// }
