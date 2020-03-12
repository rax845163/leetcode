package q764

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	have := []int{0, 0, 1, 0}
	want := 0
	assert := assert.New(t)
	assert.Equal(want, minCostClimbingStairs(have))
}
