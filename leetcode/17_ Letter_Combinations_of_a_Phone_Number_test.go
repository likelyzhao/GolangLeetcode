package leecode

import "testing"

/*
Given a string containing digits from 2-9 inclusive,
return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below.
Note that 1 does not map to any letters.

Example 1:

Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
Example 2:

Input: digits = ""
Output: []
Example 3:

Input: digits = "2"
Output: ["a","b","c"]


Constraints:

0 <= digits.length <= 4
digits[i] is a digit in the range ['2', '9'].

*/

func stringslicecompare(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Test17(t *testing.T) {

	if ans := letterCombinations17("23"); !stringslicecompare(ans, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}) {
		t.Errorf("\"23\" expected be [\"ad\", \"ae\", \"af\", \"bd\", \"be\", \"bf\", \"cd\", \"ce\", \"cf\"], but %v got", ans)
	}

	if ans := letterCombinations17(""); !stringslicecompare(ans, []string{}) {
		t.Errorf("\"\" expected be [], but %v got", ans)
	}

	if ans := letterCombinations17("2"); !stringslicecompare(ans, []string{"a", "b", "c"}) {
		t.Errorf("\"2\" expected be [\"a\", \"b\", \"c\"], but %v got", ans)
	}

}
