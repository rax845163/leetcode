package q141

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	} else if head.Next == nil || head.Next.Next == nil {
		return false
	}

	t := head.Next
	h := head.Next.Next
	for t != h && h.Next != nil && h.Next.Next != nil {
		t = t.Next
		h = h.Next.Next
	}
	return t == h
}
