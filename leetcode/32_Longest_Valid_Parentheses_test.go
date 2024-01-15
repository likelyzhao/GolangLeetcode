package leecode

import "testing"

/*
Given a string containing just the characters '(' and ')', return the length of the longest valid (well-formed) parentheses
substring
.



Example 1:

Input: s = "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()".
Example 2:

Input: s = ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()".
Example 3:

Input: s = ""
Output: 0


Constraints:

0 <= s.length <= 3 * 104
s[i] is '(', or ')'.
*/

func Test32(t *testing.T) {

	if ans := longestValidParentheses("(()"); ans != 2 {
		t.Errorf("\"(()\" expected be 2 but %v got", ans)
	}
	if ans := longestValidParentheses(")()())"); ans != 4 {
		t.Errorf("\")()())\" expected be 4 but %v got", ans)
	}
	if ans := longestValidParentheses(""); ans != 0 {
		t.Errorf("\"\" expected be 0 but %v got", ans)
	}
	if ans := longestValidParentheses("()(()"); ans != 2 {
		t.Errorf("\"()(()\" expected be 2 but %v got", ans)
	}
	if ans := longestValidParentheses("(()(((()"); ans != 2 {
		t.Errorf("\"(()(((()\" expected be 2 but %v got", ans)
	}
	if ans := longestValidParentheses("(()())"); ans != 6 {
		t.Errorf("\"(()())\" expected be 6 but %v got", ans)
	}
	if ans := longestValidParentheses("()(())"); ans != 6 {
		t.Errorf("\"()(())\" expected be 6 but %v got", ans)
	}
	if ans := longestValidParentheses("()()"); ans != 4 {
		t.Errorf("\"()()\" expected be 4 but %v got", ans)
	}
	if ans := longestValidParentheses(")(((((()())()()))()(()))("); ans != 22 {
		t.Errorf("\")(((((()())()()))()(()))(\" expected be 22 but %v got", ans)
	}

}
