package leecode

import "fmt"

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

func search33(nums []int, target int) int {
	//fmt.Println("nums", nums)
	index := -1
	start := 0
	end := len(nums) - 1

	for {
		if target == nums[start] {
			return start
		}
		if target == nums[end] {
			return end
		}

		mid := start + (end-start)/2
		//fmt.Println("start", start, mid, end, target)
		//fmt.Println(nums[start : end+1])
		if nums[mid] == target {
			index = mid
			fmt.Println(index)
			break
		}
		tempstart := start
		tempend := end

		if nums[start] < nums[mid] && nums[mid] <= nums[end] {
			//fmt.Println("start min")
			if nums[start] <= target && target <= nums[mid] {
				tempend = mid - 1
			}
			if target > nums[mid] {
				tempstart = mid + 1
			}
			if target < nums[start] || target > nums[end] {
				break
			}
		}

		if nums[start] >= nums[mid] && nums[mid] <= nums[end] {
			//fmt.Println("mid min")
			if target > nums[mid] && target <= nums[end] {
				tempstart = mid + 1
			}
			if target >= nums[start] || target < nums[mid] {
				tempend = mid - 1
			}

			if target > nums[end] && target < nums[start] {
				break
			}
		}

		if nums[start] > nums[end] && nums[mid] >= nums[start] {
			//fmt.Println("end min")
			if target <= nums[end] || target > nums[mid] {
				tempstart = mid + 1
			}

			if target >= nums[start] && target < nums[mid] {
				tempend = mid - 1
			}
			if target < nums[start] && target > nums[end] {
				break
			}
		}
		//fmt.Println(tempstart, tempend)
		if tempstart > tempend {
			break
		} else {
			start = tempstart
			end = tempend
		}

		//fmt.Println("after", start, mid, end, target)

	}

	return index

}
