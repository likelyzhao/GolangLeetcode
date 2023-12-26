package leecode

import (
	"math"
	"sort"
)

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

// AC beat 90% plus
func ThressSumClosest16(nums []int, target int) int {
	var res int
	sort.Ints(nums)
	res = nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > target && nums[i] > 0 {
			return res
		}
		start := i + 1
		end := len(nums) - 1
		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			//fmt.Println(nums[i], nums[start], nums[end], sum)
			if sum == target {
				return sum
			}
			if math.Abs(float64(sum-target)) < math.Abs(float64(res-target)) {
				res = sum
				start++
			} else if sum < target {
				start++
			} else if sum > target {
				end--
			}
		}
	}

	return res

}
