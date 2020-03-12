package q108

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		return &TreeNode{
			Val:   nums[0],
			Left:  nil,
			Right: nil,
		}
	}

	mid := len(nums) / 2
	left := nums[0:mid]
	right := []int{}
	if mid+1 <= len(nums)-1 {
		right = nums[mid+1:]
	}

	r := &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(left),
		Right: sortedArrayToBST(right),
	}
	return r
}
