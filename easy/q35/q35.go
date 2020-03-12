package q35

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	insertTo := 0
	for j, num := range nums {
		if num < target {
			insertTo = j + 1
		} else if num == target {
			return j
		} else if num > target {
			return j
		}
	}
	return insertTo
}
