package leecode

import "testing"

/*
Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.



Example 1:

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
Example 2:

Input: nums = [0,1]
Output: [[0,1],[1,0]]
Example 3:

Input: nums = [1]
Output: [[1]]


Constraints:

1 <= nums.length <= 6
-10 <= nums[i] <= 10
All the integers of nums are unique.
*/

func Test46(t *testing.T) {

	if ans := permute([]int{1, 2, 3}); !twoDslicecompare(ans, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3},
		{2, 3, 1}, {3, 1, 2}, {3, 2, 1}}) {
		t.Errorf("[1, 2, 3]   expected be [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]], but %v got", ans)
	}
	if ans := permute([]int{0, 1}); !twoDslicecompare(ans, [][]int{{0, 1}, {1, 0}}) {
		t.Errorf("[0,1]   expected be  [[0,1],[1,0]], but %v got", ans)
	}
	if ans := permute([]int{1}); !twoDslicecompare(ans, [][]int{{1}}) {
		t.Errorf("[1]   expected be  [[1]], but %v got", ans)
	}

}
