package leecode

import "testing"

/*
There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.

Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
Example 2:

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
Example 3:

Input: nums = [1], target = 0
Output: -1

Constraints:

1 <= nums.length <= 5000
-104 <= nums[i] <= 104
All values of nums are unique.
nums is an ascending array that is possibly rotated.
-104 <= target <= 104
*/
func Test33(t *testing.T) {

	if ans := search33([]int{3, 4, 5, 6, 7, 8, 1, 2}, 7); ans != 4 {
		t.Errorf("{	3,4,5,6,7,8,1,2},7  expected be 4 but %v got", ans)
	}

	if ans := search33([]int{3, 4, 5, 1, 2}, 4); ans != 1 {
		t.Errorf("{	3, 4, 5, 1, 2,} 4, expected be 1 but %v got", ans)
	}
	if ans := search33([]int{5, 1, 2, 3, 4}, 1); ans != 1 {
		t.Errorf("{	5, 1, 2, 3, 4}, 1, expected be 1 but %v got", ans)
	}

	if ans := search33([]int{4, 5, 6, 7, 8, 1, 2, 3}, 8); ans != 4 {
		t.Errorf("{	4,5,6,7,8,1,2,3}, 4, expected be 4 but %v got", ans)
	}

	if ans := search33([]int{8, 1, 2, 3, 4, 5, 6, 7}, 6); ans != 6 {
		t.Errorf("{8, 1, 2, 3, 4, 5, 6, 7}, 6, expected be 6 but %v got", ans)
	}
	if ans := search33([]int{5, 1, 3}, 4); ans != -1 {
		t.Errorf("{1}, 0, expected be -1 but %v got", ans)
	}
	if ans := search33([]int{5, 1, 3}, 0); ans != -1 {
		t.Errorf("{1}, 0, expected be -1 but %v got", ans)
	}
	if ans := search33([]int{1, 3}, 3); ans != 1 {
		t.Errorf("{1}, 0, expected be -1 but %v got", ans)
	}
	if ans := search33([]int{1, 3}, 4); ans != -1 {
		t.Errorf("{1}, 0, expected be -1 but %v got", ans)
	}
	if ans := search33([]int{4, 5, 6, 7, 0, 1, 2}, 0); ans != 4 {
		t.Errorf("{4, 5, 6, 7, 0, 1, 2}, 0, expected be 4 but %v got", ans)
	}
	if ans := search33([]int{4, 5, 6, 7, 0, 1, 2}, 3); ans != -1 {
		t.Errorf("{4, 5, 6, 7, 0, 1, 2}, 3, expected be -1 but %v got", ans)
	}
	if ans := search33([]int{1}, 0); ans != -1 {
		t.Errorf("{1}, 0, expected be -1 but %v got", ans)
	}
	if ans := search33([]int{1}, 2); ans != -1 {
		t.Errorf("{1}, 0, expected be -1 but %v got", ans)
	}

}
