package q7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	assert := assert.New(t)
	have := -2147483412
	want := -2143847412
	actual := reverse(have)
	assert.Equal(want, actual)
}
