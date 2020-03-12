package q27

func removeElement(nums []int, val int) int {
	compareToLen := len(nums)
	newLen := 0
	for i := 0; i < compareToLen; {
		if nums[i] != val {
			newLen++
			i++
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			nums[j-1] = nums[j]
		}
		compareToLen--
	}

	return newLen
}
