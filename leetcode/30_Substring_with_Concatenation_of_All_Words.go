package leecode

/*
You are given a string s and an array of strings words. All the strings of words are of the same length.

A concatenated substring in s is a substring that contains all the strings of any permutation of words concatenated.

For example, if words = ["ab","cd","ef"], then "abcdef", "abefcd", "cdabef", "cdefab", "efabcd", and "efcdab" are all concatenated strings. "acdbef" is not a concatenated substring because it is not the concatenation of any permutation of words.
Return the starting indices of all the concatenated substrings in s. You can return the answer in any order.

Example 1:

Input: s = "barfoothefoobarman", words = ["foo","bar"]
Output: [0,9]
Explanation: Since words.length == 2 and words[i].length == 3, the concatenated substring has to be of length 6.
The substring starting at 0 is "barfoo". It is the concatenation of ["bar","foo"] which is a permutation of words.
The substring starting at 9 is "foobar". It is the concatenation of ["foo","bar"] which is a permutation of words.
The output order does not matter. Returning [9,0] is fine too.
Example 2:

Input: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
Output: []
Explanation: Since words.length == 4 and words[i].length == 4, the concatenated substring has to be of length 16.
There is no substring of length 16 in s that is equal to the concatenation of any permutation of words.
We return an empty array.
Example 3:

Input: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
Output: [6,9,12]
Explanation: Since words.length == 3 and words[i].length == 3, the concatenated substring has to be of length 9.
The substring starting at 6 is "foobarthe". It is the concatenation of ["foo","bar","the"] which is a permutation of words.
The substring starting at 9 is "barthefoo". It is the concatenation of ["bar","the","foo"] which is a permutation of words.
The substring starting at 12 is "thefoobar". It is the concatenation of ["the","foo","bar"] which is a permutation of words.

Constraints:

1 <= s.length <= 104
1 <= words.length <= 5000
1 <= words[i].length <= 30
s and words[i] consist of lowercase English letters.
*/
func removewords(word string, words map[string]int) map[string]int {
	words[word]--
	return words
}
func countremainword(words map[string]int) int {
	res := 0
	for _, v := range words {
		res += v
	}
	return res
}

func findSubstring(s string, words []string) []int {
	var res []int
	wordlen := len(words[0])
	if len(s) < len(words)*wordlen {
		return []int{}
	}
	//totalwords := len(words)
	wordmap := map[string]int{}
	for _, w := range words {
		wordmap[w]++
	}
	if len(wordmap) == 1 && wordlen == 1 {
		if len(s) >= wordmap[words[0]] {
			for i := 0; i < len(s); i++ {
				if i+wordmap[words[0]] > len(s) {
					//fmt.Println(wordlen)
					break
				}
				res = append(res, i)

			}
		}
		return res
	}

	for i := 0; i < len(s); i++ {
		if i+len(words)*wordlen > len(s) {
			//fmt.Println(res)
			return res
		}
		//deepcopy
		remain := make(map[string]int, len(wordmap))
		for w, count := range wordmap {
			remain[w] = count
		}
		tempi := i
		for {
			matchflag := false
			for w, count := range remain {
				if count <= 0 {
					continue
				}
				if s[tempi:tempi+wordlen] == w {
					remain = removewords(w, remain)
					//fmt.Println("remain after", remain)
					tempi += wordlen
					//fmt.Println("match", w)
					matchflag = true
				}
			}
			//fmt.Println("tempi", tempi)
			//fmt.Println(s[tempi:tempi+wordlen] == remain[0])
			//fmt.Println(remain)

			if countremainword(remain) == 0 {
				res = append(res, i)
				//fmt.Println(res)
				break
			}
			if tempi+countremainword(remain)*wordlen > len(s) {
				//fmt.Println(res)
				return res
			}
			if !matchflag {
				break
			}
		}

	}
	//fmt.Println(res)
	return res

}
