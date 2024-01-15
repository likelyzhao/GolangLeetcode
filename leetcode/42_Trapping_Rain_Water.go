package leecode

/*
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

Example 1:

Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.
Example 2:

Input: height = [4,2,0,3,2,5]
Output: 9

Constraints:

n == height.length
1 <= n <= 2 * 104
0 <= height[i] <= 105
*/

func findupbound(height []int) []int {
	var res []int
	var upflag bool
	lastheight := height[0]
	cadidateidx := 0
	for idx, v := range height {
		if v > lastheight {
			cadidateidx = idx
			upflag = true
		}
		if v < lastheight && upflag {
			res = append(res, cadidateidx)
			upflag = false
		}
		lastheight = v
	}

	return res
}

func calcuwater(height []int, boundloc []int, idx int, dirction int, res *int) {
	if dirction == 0 {
		nextidx := boundloc[0]
		nextvalue := height[nextidx]
		// get max

		for _, v := range boundloc {
			if v < idx && height[v] > nextvalue {
				nextvalue = height[v]
				nextidx = v
			}
		}

		//fmt.Println("sum", nextidx, idx)
		minheight := min(height[nextidx], height[idx])
		for i := nextidx + 1; i < idx; i++ {
			if minheight-height[i] > 0 {
				*res += minheight - height[i]
			}
		}
		//fmt.Println("res", *res)
		if nextidx > 0 {
			calcuwater(height, boundloc, nextidx, 0, res)
		}

	}
	if dirction == 1 {
		nextidx := boundloc[len(boundloc)-1]
		nextvalue := height[nextidx]
		for _, v := range boundloc {
			if v > idx && height[v] > nextvalue {
				nextidx = v
				nextvalue = height[v]
			}
		}

		//fmt.Println(dirction, "sum", nextidx, idx)
		minheight := min(height[nextidx], height[idx])
		for i := idx + 1; i < nextidx; i++ {
			if minheight-height[i] > 0 {
				*res += minheight - height[i]
			}
		}
		//fmt.Println("res", *res)
		//fmt.Println(nextidx)
		if nextidx < len(height)-1 {
			//fmt.Println("nextidx", nextidx)
			calcuwater(height, boundloc, nextidx, 1, res)
		}

	}
}

func trap(height []int) int {
	var loc []int
	loc = append(loc, 0)
	loc = append(loc, findupbound(height)...)
	loc = append(loc, len(height)-1)

	//find maxvalue
	//fmt.Println("loc", loc)
	maxvalue := 0
	maxidx := 0
	for _, v := range loc {
		if height[v] > maxvalue {
			maxvalue = height[v]
			maxidx = v
		}
	}
	maxvalue = 0
	calcuwater(height, loc, maxidx, 0, &maxvalue)
	if maxidx != loc[len(loc)-1] {
		calcuwater(height, loc, maxidx, 1, &maxvalue)
	}
	return maxvalue
}
