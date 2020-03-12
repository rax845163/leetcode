package q112

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	return handler(root, sum)
}

func handler(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Right == nil && root.Left == nil {
		return sum == root.Val
	}
	return handler(root.Left, sum-root.Val) || handler(root.Right, sum-root.Val)
}
