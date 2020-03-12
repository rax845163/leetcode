package q67

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddBinary(t *testing.T) {
	assert := assert.New(t)
	have := []string{
		"1010",
		"1011",
	}
	want := "10101"
	actual := addBinary(have[0], have[1])
	assert.Equal(want, actual)
}
