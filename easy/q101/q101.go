package q101

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isSymmetricTree(root.Left, root.Right)
}

func isSymmetricTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	if p.Val != q.Val {
		return false
	}
	return isSymmetricTree(p.Left, q.Right) && isSymmetricTree(p.Right, q.Left)
}
