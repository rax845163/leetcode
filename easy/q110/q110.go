package q110

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	l := getDepth(root.Left, 1)
	r := getDepth(root.Right, 1)
	diff := l - r
	if diff < 0 {
		diff = -diff
	}
	return isBalanced(root.Right) && isBalanced(root.Left) && diff <= 1
}

func getDepth(root *TreeNode, dep int) int {
	if root == nil {
		return dep
	}

	l := getDepth(root.Left, dep+1)
	r := getDepth(root.Right, dep+1)
	if l > r {
		return l
	}
	return r
}
