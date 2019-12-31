package q12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	type tc struct {
		have int
		want string
	}
	tcs := []tc{
		tc{
			have: 1,
			want: "I",
		},
		tc{
			have: 3,
			want: "III",
		},
		tc{
			have: 4,
			want: "IV",
		},
		tc{
			have: 9,
			want: "IX",
		},
		tc{
			have: 58,
			want: "LVIII",
		},
		tc{
			have: 1994,
			want: "MCMXCIV",
		},
	}

	assert := assert.New(t)
	for i, tc := range tcs {
		assert.Equal(tc.want, intToRoman(tc.have), "case %d", i)
	}
}
