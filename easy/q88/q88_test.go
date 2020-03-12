package q88

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	a1 := []int{0}
	m := 0
	a2 := []int{2}
	n := 1
	assert := assert.New(t)
	want := []int{2}
	merge(a1, m, a2, n)
	assert.Equal(want, a1)
}
