package q125

import (
	"strings"
)

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	s = strings.ToLower(s)

	i, j := 0, len(s)-1
	for j-i > 0 {
		if i == j {
			return true
		}
		a, b := s[i], s[j]
		if !(('a' <= a && a <= 'z') || ('0' <= a && a <= '9')) {
			i++
			continue
		}
		if !(('a' <= b && b <= 'z') || ('0' <= b && b <= '9')) {
			j--
			continue
		}
		if a != b {
			return false
		}
		i++
		j--
	}
	return true
}
