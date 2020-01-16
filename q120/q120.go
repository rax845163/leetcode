package q120

import (
	"math"
)

func minimumTotal(triangle [][]int) int {

	if len(triangle) == 1 {
		return triangle[0][0]
	}

	for rowIndex := len(triangle) - 2; rowIndex >= 0; rowIndex-- {
		for i, e := range triangle[rowIndex] {
			left := e + triangle[rowIndex+1][i]
			right := e + triangle[rowIndex+1][i+1]
			triangle[rowIndex][i] = int(math.Min(float64(left), float64(right)))
		}
	}

	return triangle[0][0]
}
