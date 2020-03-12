package q58

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	strtok := strings.Split(s, " ")
	for i := len(strtok) - 1; i >= 0; i-- {
		if strtok[i] == "" {
			continue
		}
		return len(strtok[i])
	}

	return 0
}
