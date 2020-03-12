package q67

import (
	"fmt"
)

func addBinary(a string, b string) string {
	ret := ""
	c := uint8(0)
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		var r uint8
		if i >= 0 {
			r += (a[i] - '0')
		}
		if j >= 0 {
			r += (b[j] - '0')
		}
		r += (c - 0)
		ret = fmt.Sprintf("%d%s", r%2, ret)
		c = r / 2
	}
	if c != 0 {
		ret = fmt.Sprintf("1%s", ret)
	}
	return ret
}
