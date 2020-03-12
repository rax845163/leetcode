package q166

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	have_num := 4
	have_demin := 333
	want := ""
	assert := assert.New(t)
	actual := fractionToDecimal(have_num, have_demin)
	assert.Equal(want, actual)
}
