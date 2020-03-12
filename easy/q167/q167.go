package q167

func twoSum(numbers []int, target int) []int {
	l := 0
	h := len(numbers) - 1

	for l < h {
		sum := numbers[l] + numbers[h]
		if sum == target {
			break
		} else if sum > target {
			h--
		} else {
			l++
		}
	}

	return []int{l + 1, h + 1}
}
