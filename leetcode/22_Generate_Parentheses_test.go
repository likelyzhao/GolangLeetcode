package leecode

import (
	"testing"

	"github.com/thoas/go-funk"
)

/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

Example 1:

Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
Example 2:

Input: n = 1
Output: ["()"]

Constraints:

1 <= n <= 8
*/
func compare22res(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if !funk.ContainsString(b, a[i]) {
			return false
		}
	}
	return true
}
func Test22(t *testing.T) {

	if ans := generateParenthesis(3); !compare22res(ans, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}) {
		t.Errorf("3 expected be \"((()))\", \"(()())\", \"(())(\", \"()(())\", \"()()()\" but %v got", ans)
	}
}
