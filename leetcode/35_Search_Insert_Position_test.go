package leecode

import "testing"

/*
Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.



Example 1:

Input: nums = [1,3,5,6], target = 5
Output: 2
Example 2:

Input: nums = [1,3,5,6], target = 2
Output: 1
Example 3:

Input: nums = [1,3,5,6], target = 7
Output: 4


Constraints:

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums contains distinct values sorted in ascending order.
-104 <= target <= 104
*/

func Test35(t *testing.T) {

	if ans := searchInsert([]int{1, 3, 5, 6}, 5); ans != 2 {
		t.Errorf("{	1,2,3,5},5  expected be 2 but %v got", ans)
	}
	if ans := searchInsert([]int{1, 3, 5, 6}, 2); ans != 1 {
		t.Errorf("{	1, 3, 5, 6},2  expected be 1 but %v got", ans)
	}
	if ans := searchInsert([]int{1, 3, 5, 6}, 7); ans != 4 {
		t.Errorf("{	1, 3, 5, 6}, 7  expected be 4 but %v got", ans)
	}
}
