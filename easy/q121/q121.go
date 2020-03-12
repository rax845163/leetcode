package q121

func maxProfit(prices []int) int {
	max := 0
	low := prices[0]
	for _, p := range prices {
		if (p - low) > max {
			max = p - low
		}
		if p < low {
			low = p
		}
	}
	return max
}
