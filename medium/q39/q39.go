package q39

var dp = make(map[int][][]int)

func combinationSum(candidates []int, target int) [][]int {
	if a, ok := helper(candidates, target, 0, nil); ok {
		return a
	}
	return nil
}

func helper(candidates []int, target int, index int, m [][]int) ([][]int, bool) {
	if index >= len(candidates) {
		return nil, false
	} else if target < 0 {
		return nil, false
	}
	if v, exist := dp[target]; exist {
		return v, true
	}
	if target == candidates[index] {
		m = add(m, candidates[index])
		dp[target] = m
		return m, true
	}

	chose, chooseOk := helper(candidates, target-candidates[index], index, m)
	if chooseOk {

	}

	noChose, noChooseOk := helper(candidates, target, index+1, m)

	var ret [][]int
	if !chooseOk && !noChooseOk {
		return nil, false
	} else if chooseOk && noChooseOk {
		ret = merge(chose, noChose)
	} else if chooseOk {
		ret = chose
	} else if noChooseOk {
		ret = noChose
	}

	return ret, true
}

func add(m [][]int, a int) [][]int {
	if m == nil || len(m) == 0 {
		m = append(m, []int{a})
	} else {
		for i, v := range m {
			m[i] = append(v, a)
		}
	}
	return m
}

func merge(a, b [][]int) [][]int {
	m := [][]int{}
	for _, v := range a {
		m = append(m, v)
	}
	for _, v := range b {
		m = append(m, v)
	}
	return m
}
