package q58

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	assert := assert.New(t)
	have := "b   a"
	want := 1
	actual := lengthOfLastWord(have)
	assert.Equal(want, actual)
}
