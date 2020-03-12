package q680

func validPalindrome(s string) bool {
	return helper(s, true)
}

func helper(s string, canIgnore bool) bool {
	if s == "" {
		return true
	}
	i := 0
	j := len(s) - 1
	for j-i > 0 {
		a := s[i]
		b := s[j]
		if a == b {
			i++
			j--
			continue
		}
		if !canIgnore {
			return false
		}
		return helper(s[i:j], false) || helper(s[i+1:j+1], false)
	}
	return true
}
