package leecode

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

func longestValidParentheses(s string) int {
	var matchcadidate []byte
	resmap := make(map[int]int)
	maxres := 0
	lastmatch := true
	//laselen := -1
	for _, c := range s {
		if c == '(' {
			matchcadidate = append(matchcadidate, byte(c))
			if len(matchcadidate) > len(resmap) {
				resmap[len(matchcadidate)-1] = 0
			}
			lastmatch = false
			//fmt.Println(resmap, tempres, maxres)
			continue
		}

		if c == ')' && len(matchcadidate) > 0 {
			matchcadidate = matchcadidate[:len(matchcadidate)-1]
			resmap[len(matchcadidate)] += 2
			//fmt.Println(resmap, "check")
			if lastmatch {
				resmap[len(matchcadidate)] += resmap[len(matchcadidate)+1]
				resmap[len(matchcadidate)+1] = 0
			}
			lastmatch = true
		} else {
			//fmt.Println("here")
			matchcadidate = []byte{}
			for _, v := range resmap {
				if v > maxres {
					maxres = v
				}
			}
			resmap = make(map[int]int)
			//tempres = 0
			lastmatch = false
		}
		//fmt.Println(resmap, tempres, maxres)
	}

	for _, v := range resmap {
		if v > maxres {
			maxres = v
		}
	}

	return maxres

}
