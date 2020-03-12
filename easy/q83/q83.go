package q83

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	prevPtr := head
	findV := head.Val
	cur := prevPtr.Next
	prevPtr.Next = nil
	for cur != nil {
		curV := cur.Val
		if curV != findV {
			prevPtr.Next = cur
			prevPtr = cur
			findV = curV
			cur = cur.Next
			prevPtr.Next = nil
		} else {
			tmp := cur
			cur = cur.Next
			tmp.Next = nil
		}
	}
	return head
}
