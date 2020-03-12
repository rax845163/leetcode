package q104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return ok(root, 0)
}

func ok(root *TreeNode, dep int) int {
	if root == nil {
		return dep
	}

	l := ok(root.Left, dep+1)
	r := ok(root.Right, dep+1)
	if l > r {
		return l
	}
	return r
}
