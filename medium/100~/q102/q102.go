package q102

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	m := [][]int{}
	return handler(root, 1, m)
}
func handler(root *TreeNode, dep int, m [][]int) [][]int {
	if root == nil {
		return m
	}
	for len(m) < dep {
		m = append(m, []int{})
	}
	m[dep-1] = append(m[dep-1], []int{root.Val}...)
	m = handler(root.Left, dep+1, m)
	m = handler(root.Right, dep+1, m)
	return m
}
