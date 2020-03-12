package q111

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	r := helper(root, 1)
	if r == -1 {
		return 0
	}
	return r
}

func helper(root *TreeNode, d int) int {
	if root == nil {
		return -1
	}

	l := helper(root.Left, d+1)
	r := helper(root.Right, d+1)
	if l == -1 && r == -1 {
		return d
	} else if l == -1 {
		return r
	} else if r == -1 {
		return l
	}
	if l > r {
		return r
	}
	return l
}
