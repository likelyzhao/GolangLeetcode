package leecode

import "strings"

func checkPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	if s[0] == s[len(s)-1] {
		return checkPalindrome(s[1 : len(s)-1])
	} else {
		return false
	}

}

func growPalindrome(sori string, checkstring string, idx int) (growthstring string) {

	startpos := idx - len(checkstring)
	endpos := idx + 1
	var flag bool
	flag = true
	for i := 0; i < len(checkstring); i++ {
		if checkstring[i] != checkstring[0] {
			flag = false
			break
		}
	}
	if flag == true {
		for {
			var startflag, endflag int
			startflag = 0
			endflag = 0
			if startpos < 0 || sori[startpos] != checkstring[0] {
				startflag = 1
			}

			if endpos >= len(sori) || sori[endpos] != checkstring[0] {
				endflag = 1
			}

			if startpos >= 0 && startflag != 1 {
				startpos--
			}

			if endpos < len(sori) && endflag != 1 {
				endpos++
			}

			if startflag == 1 && endflag == 1 {
				break
			}
			//	endpos++
		}
	}

	for {
		if startpos < 0 || endpos >= len(sori) {
			break
		}
		if sori[startpos] != sori[endpos] {
			break
		}
		startpos--
		endpos++
	}

	var tempstr []byte
	for i := startpos + 1; i < endpos; i++ {
		tempstr = append(tempstr, sori[i])
	}

	return string(tempstr)

}

func LongestPalindrome(s string) string {
	var tempstring []byte
	longestPalindrome := string(s[0])
	tempstring = append(tempstring, s[0])

	for i := 1; i < len(s); i++ {
		if strings.Contains(string(tempstring), string(s[i])) {
			idx := strings.Index(string(tempstring), string(s[i]))
			checkstring := append(tempstring[idx:len(tempstring)], s[i])
			println(string(checkstring))
			if checkPalindrome(string(checkstring)) {
				temp := growPalindrome(s, string(checkstring), i)
				if len(temp) > len(longestPalindrome) {
					longestPalindrome = temp
				}
				tempstring = tempstring[idx+1 : len(tempstring)]
				tempstring = append(tempstring, s[i])
			} else {
				tempstring = tempstring[idx+1 : len(tempstring)]
				tempstring = append(tempstring, s[i])
			}
		} else {
			tempstring = append(tempstring, s[i])
		}

	}
	return longestPalindrome
}
