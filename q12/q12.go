package q12

import (
	"log"
)

// var romanNumMap = map[string]int{
// "M":  1000,
// "CM": 900,
// "D":  500,
// "CD": 400,
// "C":  100,
// "XC": 90,
// "L":  50,
// "XL": 40,
// "X":  10,
// "IX": 9,
// "V":  5,
// "IV": 4,
// "I":  1,
// }

var a = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
var b = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var c = []bool{false, true, false, true, false, true, false, true, false, true, false, true, false}

func intToRoman(num int) string {
	var res string
	// log.Println("--start--")
	for num > 0 {
		for i, n := range b {
			// log.Printf("n %d, num %d\n", n, num)
			if num/n > 0 && !c[i] {
				count := num / n
				for count > 0 {
					res += a[i]
					count--
				}
				num = num - num/n*n
			} else if c[i] && num/n == 1 {
				res += a[i]
				num = num - n
			}
		}
	}

	return res
}
