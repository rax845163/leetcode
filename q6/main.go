package main

import (
	"log"
)

func convert(s string, numRows int) string {
	n := make([][]byte, numRows)
	add := true
	row := 0
	for i := 0; i < len(s); i++ {
		n[row] = append(n[row], s[i])
		if numRows == 1 {
			continue
		}
		if add {
			row++
		} else {
			row--
		}
		if row == numRows-1 {
			add = false
		} else if row == 0 {
			add = true
		}
	}
	result := make([]byte, 0, len(s))
	for _, a := range n {
		result = append(result, a...)
	}
	return string(result)
}

func main() {
	s := "AV"
	n := 1
	log.Println(convert(s, n))
}
