package leecode

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

func premutesub(nums []int, tempans []int, ans *[][]int) {
	if len(nums) == 1 {
		tempans = append(tempans, nums[0])
		*ans = append(*ans, tempans)
		return
	}
	for i := 0; i < len(nums); i++ {
		tempansnext := append(tempans, nums[i])
		//fmt.Println("nums=", nums)
		//fmt.Println("i=", i)
		//fmt.Println("tempans", tempansnext)
		//fmt.Println("nextnums", append(nums[0:i], nums[i+1:]...))
		//fmt.Println("nums = ", nums)
		var nextnum []int
		for j := 0; j < len(nums); j++ {
			if j != i {
				nextnum = append(nextnum, nums[j])
			}
		}
		//fmt.Println("nextnums", nextnum)
		premutesub(nextnum, tempansnext, ans)
	}
}

func permute(nums []int) [][]int {
	var ans [][]int
	premutesub(nums, []int{}, &ans)
	return ans

}
