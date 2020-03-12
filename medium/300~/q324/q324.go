package q324

func wiggleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	odds := make([]int, 0, len(nums)/2+1)
	evens := make([]int, 0, len(nums)/2+1)

	if nums[0] > nums[1] {
		nums[0], nums[1] = nums[1], nums[0]
	}
	odds = append(odds, nums[1])
	evens = append(evens, nums[0])

	for i := 2; i < len(nums)-1; i++ {
		cur := nums[i]
		a := odds[len(odds)-1]
		b := evens[len(evens)-1]

		if len(odds)+len(evens)%2 == 0 {
			if cur > a {
				odds[len(odds)-1] = cur
				evens = append(evens, a)
			} else if cur < b {
				odds[len(odds)-1] = a
				evens[len(evens)-1] = cur
				odds = append(odds, b)
			} else {
				evens = append(evens, cur)
			}
		} else {
			if cur > b {
				odds = append(odds, cur)
			} else if cur < b {
				evens[len(evens)-1] = cur
				odds = append(odds, b)
			}
		}
	}

	i, j := 0, 0
	for ; i < len(odds) && j < len(evens); i, j = i+1, j+1 {
		nums[i+j] = evens[i]
		nums[i+j+1] = odds[j]
	}
	if len(odds) < len(evens) {
		nums[len(nums)-1] = evens[len(evens)-1]
	}
}
