package q122

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	have := []int{7, 1, 5, 3, 6, 4}
	want := 7
	assert := assert.New(t)
	assert.Equal(want, maxProfit(have))
}
