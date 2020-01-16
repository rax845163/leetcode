package q120

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinimumTotal(t *testing.T) {

	type testCase struct {
		have [][]int
		want int
	}

	testCases := []testCase{
		testCase{
			have: [][]int{
				[]int{2},
				[]int{3, 4},
				[]int{6, 5, 7},
				[]int{4, 1, 8, 3},
			},
			want: 11,
		},
	}

	assert := assert.New(t)
	for _, tc := range testCases {
		actual := minimumTotal(tc.have)
		assert.Equal(tc.want, actual)
	}

}
