package q203

type ListNode struct {
	Val  int
	Next *ListNode
}

// func removeElements(head *ListNode, val int) *ListNode {
// 	if head == nil {
// 		return head
// 	}
// 	var prev *ListNode
// 	cur := head
// 	for cur != nil {
// 		if cur.Val == val {
// 			if prev != nil {
// 				prev.Next = cur.Next
// 				tmp := cur
// 				cur = cur.Next
// 				tmp.Next = nil
// 				continue
// 			} else {
// 				// On head
// 				tmp := head
// 				head = head.Next
// 				tmp.Next = nil
// 				cur = cur.Next
// 				continue
// 			}
// 		}
// 		prev = cur
// 		cur = cur.Next
// 	}
// 	return head
// }

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return head
	}
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
			continue
		}
		cur = cur.Next
	}
	return head
}
