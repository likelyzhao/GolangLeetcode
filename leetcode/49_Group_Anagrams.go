package leecode

/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:

Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
Example 2:

Input: strs = [""]
Output: [[""]]
Example 3:

Input: strs = ["a"]
Output: [["a"]]

Constraints:

1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] consists of lowercase English letters.
*/
//const string = "abcdefghijklmnopqrstuvwxyz"

func calulatestrkey(str string) (key [26]byte) {
	for _, byte_s := range str {
		key[byte_s-'a']++
	}
	return key
}

func groupAnagrams(strs []string) [][]string {
	var res [][]string
	strmap := make(map[[26]byte][]string)

	for _, str := range strs {
		key := calulatestrkey(str)
		strmap[key] = append(strmap[key], str)

	}
	for _, v := range strmap {
		res = append(res, v)
		//fmt.Println(v)
	}
	return res

}
