package leecode

import "math"
import "fmt"

func IsPalindrome(x int) bool {
	var rid int
	rid = x
	var res int
	fmt.Println(math.MinInt32)
	if (x > math.MaxInt32) || (x < math.MinInt32) {
		return false
	}
	for {
		res = rid%10 + 10*res
		rid = rid / 10
		if rid == 0 {
			break
		}
	}
	//res = digits[0]

	if res == x {
		return true
	}
	return false

}
