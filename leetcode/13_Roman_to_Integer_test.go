package leecode

import "testing"

/*
Input: s = "III"
Output: 3
Explanation: III = 3.
Example 2:

Input: s = "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.
Example 3:

Input: s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
*/

func TestAdd(t *testing.T) {
	if ans := romanToInt13("III`"); ans != 3 {
		t.Errorf("III expected be 3, but %d got", ans)
	}

	if ans := romanToInt13("LVIII"); ans != 58 {
		t.Errorf("LVIII expected be 58, but %d got", ans)
	}

	if ans := romanToInt13("MCMXCIV"); ans != 1994 {
		t.Errorf("MCMXCIV expected be 1994, but %d got", ans)
	}
}
