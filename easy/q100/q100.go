package q100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func toSlice(t *TreeNode) []int {
	if t == nil {
		return []int{}
	}

	l := toSlice(t.Left)
	r := toSlice(t.Right)
	return append(l, append(r, t.Val)...)
}

// func isSameTree(p *TreeNode, q *TreeNode) bool {
// 	pS := toSlice(p)
// 	qS := toSlice(q)
// 	if len(pS) != len(qS) {
// 		return false
// 	}

// 	for i := 0; i < len(pS); i++ {
// 		if pS[i] != qS[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
