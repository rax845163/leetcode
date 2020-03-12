package q26

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	count := 1
	current := nums[0]
	for _, v := range nums {
		if current == v {
			continue
		} else {
			nums[count] = v
			current = v
			count++
		}
	}
	nums = nums[:count]
	return len(nums)
}
