package q287

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	have := []int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1}
	want := 9
	assert := assert.New(t)
	assert.Equal(want, findDuplicate(have))
}
