package q39

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	have := []int{2, 3, 6, 7}
	have_b := 7
	want := [][]int{}
	assert := assert.New(t)
	actual := combinationSum(have, have_b)
	assert.Equal(want, actual)
}
