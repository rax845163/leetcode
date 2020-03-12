package q645

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	have := []int{1, 2, 2, 4}
	want := []int{2, 3}
	assert := assert.New(t)
	assert.Equal(want, findErrorNums(have))
}
