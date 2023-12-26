package leecode

import "testing"

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.

Constraints:

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lowercase English letters.
*/
func Test14(t *testing.T) {

	if ans := LongestCommonPrefix14([]string{"flower", "flow", "flight"}); ans != "fl" {
		t.Errorf("[\"flower\",\"flow\",\"flight\"] expected be \"fl\", but %s got", ans)
	}

	if ans := LongestCommonPrefix14([]string{"dog", "racecar", "car"}); ans != "" {
		t.Errorf("[\"dog\",\"racecar\",\"car\"] expected be \"\", but %s got", ans)
	}

}
