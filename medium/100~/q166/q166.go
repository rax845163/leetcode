package q166

import (
	"fmt"
)

func fractionToDecimal(numerator int, denominator int) string {
	q := numerator / denominator
	r := numerator % denominator

	fractions := []string{}
	repeatingFractionIndex := -1
	remainToFractionIndex := map[int]int{}
	for {
		if r == 0 {
			break
		}

		// Init paddingZeroes with -1
		paddingZeroes := -1
		for r < denominator {
			paddingZeroes++
			r *= 10
		}
		f := ""
		q := r / denominator
		r %= denominator
		for paddingZeroes > 0 {
			f = fmt.Sprintf("0%s", f)
			paddingZeroes--
		}
		f = fmt.Sprintf("%s%d", f, q)
		fractions = append(fractions, f)
		if index, exist := remainToFractionIndex[r]; exist {
			repeatingFractionIndex = index
			break
		}
		remainToFractionIndex[r] = len(fractions) - 1
	}

	if len(fractions) == 0 {
		return fmt.Sprintf("%d", q)
	}

	fractionString := ""
	for i, fraction := range fractions {
		if i == repeatingFractionIndex {
			fractionString = fmt.Sprintf("%s(%s", fractionString, fraction)
		} else {
			fractionString = fmt.Sprintf("%s%s", fractionString, fraction)
		}
	}
	if repeatingFractionIndex != -1 {
		fractionString = fmt.Sprintf("%s)", fractionString)
	}

	return fmt.Sprintf("%d.%s", q, fractionString)
}
