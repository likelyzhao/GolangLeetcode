package leecode

import (
	"sort"
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

func Contains(a [][]int, b []int) bool {
	//sort.Ints(b)
	for i := 0; i < len(a); i++ {
		if tripletcomp(b, a[i]) {
			return true
		}
	}
	return false
}
func ContainsInt(a []int, b int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == b {
			return true
		}
	}
	return false
}

func tripletcomp(a, b []int) bool {

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			//fmt.Println(a, b)
			return false
		}
	}
	return true
}

func ThreeSum15(nums []int) [][]int {
	return ThreeSum15V2(nums)
}

func skippass(nums []int, start int) int {
	start++
	for start <= len(nums)-2 {
		//fmt.Println(start, nums[start], len(nums))
		if nums[start] == nums[start-1] {
			start++
		} else {
			return start
		}
	}
	return start
}

// AC
func ThreeSum15V2(nums []int) [][]int {
	sort.Ints(nums)
	//fmt.Println(nums)
	var result [][]int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		start := i + 1
		end := len(nums) - 1
		for start < end {
			if nums[i]+nums[start]+nums[end] == 0 {
				//fmt.Println(nums[i], nums[start], nums[end])
				tempints := []int{nums[i], nums[start], nums[end]}
				//if !Contains(result, tempints) {
				result = append(result, tempints)
				//}
				//start++
				start = skippass(nums, start)
			} else if nums[i]+nums[start]+nums[end] > 0 {
				end--
			} else if nums[i]+nums[start]+nums[end] < 0 {
				//start++
				start = skippass(nums, start)
			}
		}
	}

	return result

}

// overTimeLimits
func ThreeSum15V1(nums []int) [][]int {
	sort.Ints(nums)
	//fmt.Println(nums)
	var result [][]int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					tempints := []int{nums[i], nums[j], nums[k]}
					//sort.Ints(tempints)
					if !Contains(result, tempints) {
						result = append(result, tempints)
					}
				}

			}
		}
	}

	return result

}
