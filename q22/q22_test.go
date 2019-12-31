package q22

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGen(t *testing.T) {

	type tc struct {
		have int
		want []string
	}

	tcs := []tc{
		tc{
			have: 4,
			want: []string{
				"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()",
			},
		},
		tc{
			have: 3,
			want: []string{
				"((()))",
				"(()())",
				"(())()",
				"()(())",
				"()()()",
			},
		},
		tc{
			have: 10,
			want: []string{},
		},
	}

	assert := assert.New(t)
	for _, tc := range tcs {
		actual := generateParenthesis(tc.have)
		t.Logf("%+v", actual)
		for _, a := range actual {
			assert.Contains(tc.want, a)
		}
	}
}

func benchmarkGenerateParenthesis(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateParenthesis(i)
	}
}

func BenchmarkGenerateParenthesis3(b *testing.B) {
	generateParenthesis(3)
}

func BenchmarkGenerateParenthesis10(b *testing.B) {
	benchmarkGenerateParenthesis(10, b)
}
