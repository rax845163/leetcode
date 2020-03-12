package q53

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	assert := assert.New(t)
	have := []int{-1}
	want := -1
	actual := maxSubArray(have)
	assert.Equal(want, actual)
}
