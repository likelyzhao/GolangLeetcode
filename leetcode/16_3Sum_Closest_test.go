package leecode

import "testing"

/*
Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.

Return the sum of the three integers.

You may assume that each input would have exactly one solution.



Example 1:

Input: nums = [-1,2,1,-4], target = 1
Output: 2
Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
Example 2:

Input: nums = [0,0,0], target = 1
Output: 0
Explanation: The sum that is closest to the target is 0. (0 + 0 + 0 = 0).


Constraints:

3 <= nums.length <= 500
-1000 <= nums[i] <= 1000
-104 <= target <= 104
*/

func Test16(t *testing.T) {

	if ans := ThressSumClosest16([]int{-1, 2, 1, -4}, 1); ans != 2 {
		t.Errorf("[-1, 2, 1, -4] expected be 2, but %v got", ans)
	}

	if ans := ThressSumClosest16([]int{0, 0, 0}, 1); ans != 0 {
		t.Errorf("[0,0,0] expected be 0, but %v got", ans)
	}

	if ans := ThressSumClosest16([]int{-100, -98, -2, -1}, -101); ans != -101 {
		t.Errorf("[-100, -98, -2, -1] expected be -101, but %v got", ans)
	}

}
