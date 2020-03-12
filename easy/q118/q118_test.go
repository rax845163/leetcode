package q118

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerate(t *testing.T) {
	assert := assert.New(t)
	have := 5
	want := [][]int{}
	actual := generate(have)
	assert.Equal(want, actual)
}
