package leecode

import "testing"

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
func Test42(t *testing.T) {

	if ans := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}); ans != 6 {
		t.Errorf("{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}  expected be 6 but %v got", ans)
	}
	if ans := trap([]int{4, 2, 0, 3, 2, 5}); ans != 9 {
		t.Errorf("{	4,2,0,3,2,5}  expected be 9 but %v got", ans)
	}
	if ans := trap([]int{0}); ans != 0 {
		t.Errorf("{	0}  expected be 0 but %v got", ans)
	}
	if ans := trap([]int{5, 4, 1, 2}); ans != 1 {
		t.Errorf("{	5, 4, 1, 2}  expected be 1 but %v got", ans)
	}
	if ans := trap([]int{4, 3, 3, 9, 3, 0, 9, 2, 8, 3}); ans != 23 {
		t.Errorf("{	4,3,3,9,3,0,9,2,8,3}  expected be 1 but %v got", ans)
	}

}
