package leecode

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

var letterMap = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func check(idxList []int, idxMaxList []int) bool {
	for j := 0; j < len(idxList); j++ {
		if idxList[j] < idxMaxList[j]-1 {
			return false
		}
	}
	return true
}

func letterCombinations17(digits string) []string {
	var res []string
	if len(digits) == 0 {
		return []string{}
	}
	idxList := make([]int, len(digits))
	for i := 0; i < len(digits); i++ {
		idxList[i] = 0
	}
	idxMaxList := make([]int, len(digits))
	for i := 0; i < len(digits); i++ {
		idxMaxList[i] = len(letterMap[digits[len(digits)-1-i]])
	}
	//fmt.Println(idxMaxList)

	for {
		for i := 0; i < len(idxList); i++ {
			if idxList[i] >= idxMaxList[i] {
				idxList[i] = 0
				if i < len(idxList)-1 {
					idxList[i+1]++
				}
			}
		}
		var tempstr string
		for i := len(idxList) - 1; i >= 0; i-- {
			tempstr +=
				string(letterMap[digits[len(digits)-1-i]][idxList[i]])
		}
		res = append(res, tempstr)
		if check(idxList, idxMaxList) {
			break
		}

		idxList[0]++
	}

	return res
}
