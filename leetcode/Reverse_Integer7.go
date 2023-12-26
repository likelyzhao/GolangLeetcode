package leecode

import "math"

func Reverse(x int) int {
	var rid int
	rid = x
	var res int
	digits := make([]int, 0)
	for {
		digits = append(digits, rid%10)
		res = rid%10 + 10*res
		rid = rid / 10
		if rid == 0 {
			break
		}
	}
	//res = digits[0]

	if (res > math.MaxInt32) || (res < math.MinInt32) {
		return 0
	}
	return res

}


