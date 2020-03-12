package q88

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	insertIndex := n + m - 1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[insertIndex] = nums1[i]
			i--
		} else {
			nums1[insertIndex] = nums2[j]
			j--
		}
		insertIndex--
	}
	for i >= 0 {
		nums1[insertIndex] = nums1[i]
		i--
		insertIndex--
	}
	for j >= 0 {
		nums1[insertIndex] = nums2[j]
		j--
		insertIndex--
	}
}
