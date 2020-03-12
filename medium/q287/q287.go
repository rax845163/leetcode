package q287

/*
	You must not modify the array (assume the array is read only).
	You must use only constant, O(1) extra space.
	Your runtime complexity should be less than O(n2).
	There is only one duplicate number in the array, but it could be repeated more than once.
*/
func findDuplicate(nums []int) int {
	t := 0
	h := 0

	t = nums[0] - 1
	h = nums[nums[0]-1] - 1
	for t != h {
		t = nums[t] - 1
		h = nums[nums[h]-1] - 1
	}
	return nums[t]
}
