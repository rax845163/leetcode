package q83

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	n5 := &ListNode{
		Val:  3,
		Next: nil,
	}
	n4 := &ListNode{
		Val:  3,
		Next: n5,
	}
	n3 := &ListNode{
		Val:  2,
		Next: n4,
	}
	n2 := &ListNode{
		Val:  1,
		Next: n3,
	}
	n1 := &ListNode{
		Val:  1,
		Next: n2,
	}

	assert := assert.New(t)
	have := n1
	want := n1
	actual := deleteDuplicates(have)
	assert.Equal(want, actual)
}
