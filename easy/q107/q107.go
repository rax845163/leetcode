package q107

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	m := make(map[int][]int)
	insertToLevelMap(root, 0, m)

	ret := make([][]int, 0, len(m))
	for i := len(m) - 1; i >= 0; i-- {
		ret = append(ret, m[i])
	}

	return ret

}

func insertToLevelMap(root *TreeNode, dep int, m map[int][]int) {
	if root == nil {
		return
	}
	if m == nil {
		m = make(map[int][]int)
	}
	m[dep] = append(m[dep], root.Val)

	insertToLevelMap(root.Left, dep+1, m)
	insertToLevelMap(root.Right, dep+1, m)
}
