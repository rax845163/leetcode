package q168

import (
	"fmt"
)

func convertToTitle(n int) string {
	ret := ""
	i := 1
	for n > 0 {
		c := ""
		r := (n % 26)
		if r == 0 {
			c = "z"
		} else {
			c = string(r - 1 + 'A')
		}
		n = n / 26
		i++
		ret = fmt.Sprintf("%s%s", c, ret)
	}
	return ret
}
