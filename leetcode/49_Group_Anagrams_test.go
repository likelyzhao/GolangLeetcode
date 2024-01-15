package leecode

import "testing"

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

func twoDstringslicecompare(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if stringslicecompare(a[i], b[i]) == false {
			return false
		}
	}
	return true
}

func Test49(t *testing.T) {

	if ans := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}); !twoDstringslicecompare(ans, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}) {
		t.Errorf(" [\"eat\",\"tea\",\"tan\",\"ate\",\"nat\",\"bat\"]   expected be {\"bat\"}, {\"nat\", \"tan\"}, {\"ate\", \"eat\", \"tea\"}, but %v got", ans)
	}

	if ans := groupAnagrams([]string{""}); !twoDstringslicecompare(ans, [][]string{{""}}) {
		t.Errorf(" [\"\"]   expected be {\"\"}}, but %v got", ans)
	}

	if ans := groupAnagrams([]string{"a"}); !twoDstringslicecompare(ans, [][]string{{"a"}}) {
		t.Errorf(" [\"a\"]   expected be {\"a\"}}, but %v got", ans)
	}

}
