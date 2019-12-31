package main

import (
	"fmt"
	"log"
)

func lengthOfLongestSubstring(s string) int {
	log.Println("---lengthOfLongestSubstring---")
	max := 0
	start, end := 0, 0
	prevIndexMap := map[byte]int{}
	for end < len(s) {
		log.Printf("s %d, e %d, string: %s\n", start, end, s[start:end+1])
		currentLength := end - start + 1
		current := s[end]
		if _, exist := prevIndexMap[current]; !exist {
			prevIndexMap[current] = end
		} else {
			log.Println("-re-")
			tmp := start
			for tmp < prevIndexMap[current] {
				delete(prevIndexMap, s[tmp])
				tmp++
			}
			currentLength--
			start = prevIndexMap[current] + 1
			prevIndexMap[current] = end
		}
		if max < currentLength {
			max = currentLength
		}
		end++
	}

	return max
}

func main() {

	tcase := map[string]int{
		"abcabcbb": 3,
		"dvdf":     3,
		"tmmzuxt":  5,
	}

	for s, v := range tcase {
		if g := lengthOfLongestSubstring(s); g != v {
			panic(fmt.Sprintf("case %s, expect: %d, get %d", s, v, g))
		}
	}
}
