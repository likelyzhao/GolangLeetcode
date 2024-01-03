package leecode

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

func generatesub(prefix string, left []string, right []string, temp []string) []string {
	//fmt.Println(prefix)
	//fmt.Println("left", len(left), "right", len(right))
	if len(left) > len(right) {
		return temp
	}
	if len(left) == 0 {
		for i := 0; i < len(right); i++ {
			prefix += right[i]
		}
		//fmt.Println("add", prefix)

		return append(temp, prefix)
	}
	if len(left) > 0 {
		temp = generatesub(prefix+left[0], left[1:], right, temp)
	}

	if len(right) > 0 {
		temp = generatesub(prefix+right[0], left, right[1:], temp)
	}
	return temp

}

func generateParenthesis(n int) []string {
	var left, right []string
	for i := 0; i < n; i++ {
		left = append(left, "(")
		right = append(right, ")")
	}
	return generatesub("", left, right, []string{})

}
