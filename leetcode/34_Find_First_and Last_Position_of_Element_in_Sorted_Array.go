package leecode

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
func findstart(nums []int, target int, start, end int) int {
	for {
		if start >= end {
			break
		}
		mid := start + (end-start)/2
		if nums[mid] == target {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return end
}

func findend(nums []int, target int, start, end int) int {
	//fmt.Println("find end input :", nums[start:end+1])
	for {
		//fmt.Println(start, end)
		if start >= end {
			break
		}
		mid := start + (end-start)/2 + 1
		if nums[mid] == target {
			start = mid
		} else {
			end = mid - 1
		}
	}
	return end
}

func searchRange(nums []int, target int) []int {
	var res []int
	start := 0
	end := len(nums) - 1
	if len(nums) == 1 && nums[0] == target {
		return []int{0, 0}
	}
	//fmt.Println(nums)
	for end > start {
		//fmt.Println(start, end)
		if nums[start] == target {
			res = append(res, start)
			res = append(res, findend(nums, target, start, end))
			break
		} else if nums[end] == target {
			res = append(res, findstart(nums, target, start, end))
			res = append(res, end)
			break
			//return res
		} else {
			mid := start + (end-start)/2
			if nums[mid] == target {
				res = append(res, findstart(nums, target, start, mid))
				res = append(res, findend(nums, target, mid, end))
				break
			}
			if nums[mid] > target {
				end = mid - 1
			}
			if nums[mid] < target {
				start = mid + 1
			}
			//return res
		}
	}
	if len(res) == 0 {
		res = append(res, -1)
		res = append(res, -1)
	}

	return res
}
