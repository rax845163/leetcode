package q38

import (
	"fmt"
)

func countAndSay(n int) string {
	if n == 0 {
		return ""
	}

	return fmt.Sprintf("%s%s", countAndSay(n/10))
}
