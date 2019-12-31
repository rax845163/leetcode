package q22

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	children := generateParenthesis(n - 1)
	m := map[string]struct{}{}
	result := make([]string, 0, len(children)*2)
	for _, child := range children {
		post := fmt.Sprintf(`()%s`, child)
		m[post] = struct{}{}
		suffix := fmt.Sprintf(`%s()`, child)
		m[suffix] = struct{}{}

		s := 0
		for i, c := range child {
			switch c {
			case '(':
				s++
			case ')':
				s--
			}
			if s == 0 {
				wrap := fmt.Sprintf(`(%s)%s`, child[:i], child[i:])
				m[wrap] = struct{}{}
			}
		}
	}
	for k := range m {
		result = append(result, k)
	}
	return result
}
