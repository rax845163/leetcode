package q53

import (
	"math"
)

// func maxSubArray(nums []int) int {

// 	prevMax := math.MinInt32
// 	max := math.MinInt32
// 	for _, num := range nums {
// 		if prevMax+num >0{
// 			prevMax = prevMax+num
// 		}else if
// 	}
// 	return max
// }

func maxSubArray(nums []int) int {
	_, max := findMax(nums, math.MinInt32)
	return max
}

func findMax(nums []int, max int) (int, int) {
	if len(nums) == 0 {
		return 0, 0
	} else if len(nums) == 1 {
		return nums[0], nums[0]
	}

	a := nums[0]
	b, max := findMax(nums[1:], math.MinInt32)
	if a > a+b {
		return a, maxInt(a, max)
	} else if a+b > 0 {
		return a + b, maxInt(a+b, max)
	} else {
		return 0, max
	}
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
