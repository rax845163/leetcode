package q764

func minCostClimbingStairs(cost []int) int {
	minCost := make(map[int]int)
	min := helper(cost, 0, 0, minCost)
	a := helper(cost, 1, 0, minCost)
	if a < min {
		min = a
	}
	return min
}

func helper(cost []int, step int, count int, minCost map[int]int) int {
	if step >= len(cost) {
		return count
	}

	if v, exist := minCost[step]; exist {
		return v
	}

	cur := cost[step]
	one := helper(cost, step+1, count+cur, minCost)
	two := helper(cost, step+2, count+cur, minCost)
	min := one
	if min > two {
		min = two
	}
	minCost[step] = min
	return min
}
