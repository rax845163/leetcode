package q15

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	dp := make(map[int]map[int]struct{})
	empty := struct{}{}
	for i, num := range nums {
		indicesSet := dp[num]
		if indicesSet == nil {
			indicesSet = make(map[int]struct{})
		}
		indicesSet[i] = empty
		dp[num] = indicesSet
	}

	retSet := make(map[string][]int)
	for i, nums1 := range nums {
		for j, nums2 := range nums {
			if i == j {
				continue
			}
			target := 0 - (nums1 + nums2)
			if indicesSet, exist := dp[target]; exist {
				targetIndices := isRemain(indicesSet, []int{i, j})
				for _, k := range targetIndices {
					target := nums[k]
					id := getID(nums1, nums2, target)
					retSet[id] = []int{nums1, nums2, target}
				}
			}
		}
	}

	ret := make([][]int, 0, len(retSet))
	for _, nums := range retSet {
		ret = append(ret, nums)
	}
	return ret
}

func isRemain(indicesSet map[int]struct{}, toExclude []int) []int {
	tmp := make(map[int]struct{}, len(indicesSet))
	empty := struct{}{}
	for index := range indicesSet {
		tmp[index] = empty
	}

	for _, num := range toExclude {
		delete(tmp, num)
	}

	ret := make([]int, 0, len(tmp))
	for i := range tmp {
		ret = append(ret, i)
	}
	return ret
}

func getID(a, b, c int) string {
	tmp := []int{a, b, c}
	sort.Ints(tmp)
	return fmt.Sprintf("%d/%d/%d", tmp[0], tmp[1], tmp[2])
}
