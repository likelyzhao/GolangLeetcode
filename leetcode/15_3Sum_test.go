package leecode

import (
	"testing"
)

/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.

Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.
Example 2:

Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.
Example 3:

Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.

Constraints:

3 <= nums.length <= 3000
-105 <= nums[i] <= 105
*/

func twoDslicecompare(a, b [][]int) bool {
	//fmt.Println(a, b)
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if !slicecompare(a[i], b[i]) {
			//fmt.Println(a, b)
			return false
		}
	}
	return true
}

func slicecompare(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Test15(t *testing.T) {

	if ans := ThreeSum15([]int{-1, 0, 1, 2, -1, -4}); !twoDslicecompare(ans, [][]int{{-1, -1, 2}, {-1, 0, 1}}) {
		t.Errorf("[-1, 0, 1, 2, -1, -4] expected be [{-1, -1, 2}, {-1, 0, 1}], but %v got", ans)
	}

	if ans := ThreeSum15([]int{0, 1, 1}); !twoDslicecompare(ans, [][]int{}) {
		t.Errorf("[0,1,1] expected be [], but %v got", ans)
	}

	if ans := ThreeSum15([]int{0, 0, 0}); !twoDslicecompare(ans, [][]int{{0, 0, 0}}) {
		t.Errorf("[0,0,0] expected be [{0, 0, 0}], but %v got", ans)
	}

	if ans := ThreeSum15([]int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}); !twoDslicecompare(ans, [][]int{{-5, 1, 4}, {-4, 0, 4}, {-4, 1, 3}, {-2, -2, 4}, {-2, 1, 1}, {0, 0, 0}}) {
		t.Errorf("[0,0,0] expected be [[-5 1 4] [-4 0 4] [-4 1 3] [-2 -2 4] [-2 1 1] [0 0 0]], but %v got", ans)
	}

}
