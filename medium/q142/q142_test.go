package q142

import (
	"testing"
)

func Test(t *testing.T) {
	n2 := &ListNode{}
	n4 := &ListNode{
		Val:  -4,
		Next: n2,
	}
	n3 := &ListNode{
		Val:  0,
		Next: n4,
	}
	n1 := &ListNode{
		Val:  3,
		Next: n2,
	}
	n2.Val = 2
	n2.Next = n3
	detectCycle(n1)
}
