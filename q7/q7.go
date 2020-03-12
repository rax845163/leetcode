package q7

import (
	"fmt"
	"strconv"
	"strings"
)

func reverse(x int) int {
	str := strconv.FormatInt(int64(x), 10)
	rev := ""
	for i := len(str) - 1; i >= 0; i-- {
		rev = fmt.Sprintf("%s%s", rev, string(str[i]))
	}
	rev = strings.TrimSuffix(rev, "-")

	max := "2147483647"
	min := "2147483648"

	if len(rev) > len(max) {
		return 0
	} else if len(rev) < len(max) {
		ret, _ := strconv.ParseInt(rev, 10, 32)
		if x < 0 {
			ret = -ret
		}
		return int(ret)
	} else {
		comp := ""
		if x >= 0 {
			comp = max
		} else {
			comp = min
		}

		valid := true
		for i := 0; i < len(comp); i++ {
			r, _ := strconv.ParseInt(string(rev[i]), 10, 32)
			m, _ := strconv.ParseInt(string(comp[i]), 10, 32)
			if r == m {
				continue
			} else if r < m {
				break
			} else if r > m {
				valid = false
				break
			}
		}
		if !valid {
			return 0
		}

		ret, _ := strconv.ParseInt(rev, 10, 32)
		if x < 0 {
			ret = -ret
		}
		return int(ret)
	}
}
