package q9

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	rev := ""
	for i := len(str) - 1; i >= 0; i-- {
		rev = fmt.Sprintf("%s%s", rev, string(str[i]))
	}
	r, _ := strconv.ParseInt(rev, 10, 64)
	if int(r) != x {
		return false
	}
	return true
}
