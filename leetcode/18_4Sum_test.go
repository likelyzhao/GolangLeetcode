package leecode

import "testing"

/*
Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:

0 <= a, b, c, d < n
a, b, c, and d are distinct.
nums[a] + nums[b] + nums[c] + nums[d] == target
You may return the answer in any order.



Example 1:

Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
Example 2:

Input: nums = [2,2,2,2,2], target = 8
Output: [[2,2,2,2]]


Constraints:

1 <= nums.length <= 200
-109 <= nums[i] <= 109
-109 <= target <= 109
*/

func Test18(t *testing.T) {

	if ans := fourSum([]int{1, 0, -1, 0, -2, 2}, 0); !twoDslicecompare(ans,
		[][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}) {
		t.Errorf("[1, 0, -1, 0, -2, 2] expected be [{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}], but %v got", ans)
	}

	if ans := fourSum([]int{2, 2, 2, 2, 2}, 8); !twoDslicecompare(ans, [][]int{{2, 2, 2, 2}}) {
		t.Errorf("[2,2,2,2,2] expected be [2,2,2,2], but %v got", ans)
	}
	if ans := fourSum([]int{1, -2, -5, -4, -3, 3, 3, 5}, -11); !twoDslicecompare(ans, [][]int{{-5, -4, -3, 1}}) {
		t.Errorf("[1, -2, -5, -4, -3, 3, 3, 5] expected be [-5, -4, -3, 1], but %v got", ans)
	}

}
