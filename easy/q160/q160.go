package q160

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a := headA
	b := headB
	aIsSwitch := false
	bIsSwitch := false
	for a != b {
		if a.Next == nil {
			if aIsSwitch {
				return nil
			}
			aIsSwitch = true
			a = headB
		} else {
			a = a.Next
		}
		if b.Next == nil {
			if bIsSwitch {
				return nil
			}
			bIsSwitch = true
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}
