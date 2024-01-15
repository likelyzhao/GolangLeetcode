package leecode

import "testing"

/*
Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.

Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
Example 3:

Input: nums = [], target = 0
Output: [-1,-1]

Constraints:

0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums is a non-decreasing array.
-109 <= target <= 109
*/

func Test34(t *testing.T) {
	if ans := searchRange([]int{5, 7, 7, 8, 8, 10}, 8); !slicecompare(ans, []int{3, 4}) {
		t.Errorf("{	5, 7, 7, 8, 8, 10},7  expected be [3, 4] but %v got", ans)
	}
	if ans := searchRange([]int{5, 7, 7, 8, 8, 10}, 6); !slicecompare(ans, []int{-1, -1}) {
		t.Errorf("{	5, 7, 7, 8, 8, 10}, 6  expected be [-1, -1] but %v got", ans)
	}
	if ans := searchRange([]int{1}, 1); !slicecompare(ans, []int{0, 0}) {
		t.Errorf("{	1}, 1  expected be [0, 0] but %v got", ans)
	}
}
