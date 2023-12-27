package leecode

import (
	"sort"
)

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

func ThreeSum(nums []int, target int) [][]int {
	//sort.Ints(nums)
	//fmt.Println(nums)
	var result [][]int
	for i := 0; i < len(nums); i++ {
		if nums[i] > target && target > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		start := i + 1
		end := len(nums) - 1
		for start < end {
			if nums[i]+nums[start]+nums[end] == target {
				//fmt.Println(nums[i], nums[start], nums[end])
				tempints := []int{nums[i], nums[start], nums[end]}
				//if !Contains(result, tempints) {
				result = append(result, tempints)
				//}
				//start++
				start = skippass(nums, start)
			} else if nums[i]+nums[start]+nums[end] > target {
				end--
			} else if nums[i]+nums[start]+nums[end] < target {
				//start++
				start = skippass(nums, start)
			}
		}
	}

	return result

}

func fourSum(nums []int, target int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > target && target > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		newtarget := target - nums[i]
		res := ThreeSum(nums[i+1:], newtarget)
		for _, v := range res {
			v = append([]int{nums[i]}, v...)
			result = append(result, v)
		}

	}

	return result

}
