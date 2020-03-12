package q15

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	have := []int{-1, 0, 1, 2, -1, -4}
	want := [][]int{}
	assert := assert.New(t)
	actual := threeSum(have)
	assert.Equal(want, actual)
}
