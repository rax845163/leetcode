package q102

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	head := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	have := head
	want := [][]int{
		[]int{3},
		[]int{9, 20},
		[]int{15, 7},
	}
	assert := assert.New(t)
	assert.Equal(want, levelOrder(have))
}
