package q142

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	t := head
	h := head
	if t.Next == nil || t.Next.Next == nil {
		return nil
	}

	t = t.Next
	h = h.Next.Next
	for t != h && h.Next != nil && h.Next.Next != nil {
		t = t.Next
		h = h.Next.Next
	}
	if t != h {
		return nil
	}

	t = head
	for t != h {
		t = t.Next
		h = h.Next
	}
	return t
}
