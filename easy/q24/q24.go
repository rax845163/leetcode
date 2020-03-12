package q24

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	left := head
	right := head.Next
	left.Next = right.Next
	right.Next = left
	head = right

	prev := head.Next
	current := prev.Next
	for current != nil {
		if current.Next == nil {
			break
		}
		left := current
		right := current.Next
		left.Next = right.Next
		right.Next = left
		prev.Next = right
		current = left.Next
	}
	return head
}
