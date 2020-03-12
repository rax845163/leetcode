package q13

var romanToIntMap = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

func romanToInt(s string) int {
	ret := 0

	for i := 0; i < len(s); i++ {
		if i+2 <= len(s) {
			c := string(s[i : i+2])
			if v, exist := romanToIntMap[c]; exist {
				ret += v
				i++
				continue
			}
		}
		c := string(s[i])
		v := romanToIntMap[c]
		ret += v
	}

	return ret
}
