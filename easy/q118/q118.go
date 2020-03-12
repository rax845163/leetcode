package q118

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	r := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		c := i + 1
		tmp := make([]int, c)
		tmp[0] = 1
		tmp[len(tmp)-1] = 1
		for j := 1; j < i; j++ {
			tmp[j] = r[i-1][j-1] + r[i-1][j]
		}
		r[i] = tmp
	}
	return r
}
