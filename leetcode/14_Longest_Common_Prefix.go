package leecode

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
func LongestCommonPrefix14(strs []string) string {
	commonprefix := ""
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			// 逐个匹配字符串
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return commonprefix
			}
		}
		commonprefix += string(strs[0][i])
		//fmt.Println(commonprefix)
	}

	return commonprefix
}
