package q14

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	lcp := ""
	minLen := len(strs[0])
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}

	for i := 0; i < minLen; i++ {
		c := strs[0][i]
		allSame := true
		for _, str := range strs {
			if str[i] != c {
				allSame = false
				break
			}
		}
		if !allSame {
			return lcp
		}
		lcp = fmt.Sprintf("%s%s", string(lcp), string(c))
	}

	return lcp
}
