package leecode

import "testing"

/*
Given an integer array nums, find the
subarray
 with the largest sum, and return its sum.



Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.
Example 2:

Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.
Example 3:

Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.


Constraints:

1 <= nums.length <= 105
-104 <= nums[i] <= 104


Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
*/

func Test53(t *testing.T) {

	if ans := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}); ans != 6 {
		t.Errorf("{-2, 1, -3, 4, -1, 2, 1, -5, 4}  expected be 6 but %v got", ans)
	}

	if ans := maxSubArray([]int{1}); ans != 1 {
		t.Errorf("{1}  expected be 1 but %v got", ans)
	}

	if ans := maxSubArray([]int{5, 4, -1, 7, 8}); ans != 23 {
		t.Errorf("{5,4,-1,7,8}  expected be 23 but %v got", ans)
	}
	if ans := maxSubArray([]int{-1, -2}); ans != -1 {
		t.Errorf("{-1,-2}  expected be -1 but %v got", ans)
	}
	if ans := maxSubArray([]int{8, -19, 5, -4, 20}); ans != 21 {
		t.Errorf("{-1,-2}  expected be -1 but %v got", ans)
	}
	if ans := maxSubArray([]int{31, -41, 59, 26, -53, 58, 97, -93, -23, 84}); ans != 187 {
		t.Errorf("{-1,-2}  expected be -1 but %v got", ans)
	}
	if ans := maxSubArray([]int{9, 0, -2, -2, -3, -4, 0, 1, -4, 5, -8, 7, -3, 7, -6, -4, -7, -8}); ans != 11 {
		t.Errorf("{-1,-2}  expected be -1 but %v got", ans)
	}
	if ans := maxSubArray([]int{-1, 0, -2, 2}); ans != 2 {
		t.Errorf("{-1, 0, -2, 2}  expected be 2 but %v got", ans)
	}

}
