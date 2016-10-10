package leecode

import "strings"

func LengthOfLongestSubstring(s string) int {
	var substring []byte

	if len(s) == 0 {
		return 0
	}
	substring = append(substring, s[0])
	k := 1
	for i := 1; i < len(s); i++ {
		str := string(substring)
		if strings.Contains(str, string(s[i])) == true {
			idx := strings.Index(str, string(s[i]))
			if k < len(substring) {
				k = len(substring)
			}
			substring = substring[idx+1 : len(substring)]
			substring = append(substring, s[i])
		} else {
			substring = append(substring, s[i])

		}

	}
	if k < len(substring) {
		k = len(substring)
	}
	return k
}
