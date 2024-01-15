package leecode

import "fmt"

/*
You are given a 0-indexed array of integers nums of length n. You are initially positioned at nums[0].

Each element nums[i] represents the maximum length of a forward jump from index i. In other words, if you are at nums[i], you can jump to any nums[i + j] where:

0 <= j <= nums[i] and
i + j < n
Return the minimum number of jumps to reach nums[n - 1]. The test cases are generated such that you can reach nums[n - 1].



Example 1:

Input: nums = [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.
Example 2:

Input: nums = [2,3,0,1,4]
Output: 2


Constraints:

1 <= nums.length <= 104
0 <= nums[i] <= 1000
It's guaranteed that you can reach nums[n - 1].
*/

func findmin(nums []int, now int, step int, minstep *int) {
	if now+nums[now] >= len(nums)-1 {
		step += 1
		//fmt.Println(now, step)
		if *minstep > step {
			*minstep = step
			fmt.Println("minstep", *minstep)
		}
	} else {
		if step+1 >= *minstep {
			return
		} else {
			tempmax := nums[now]
			for i := nums[now]; i > 0; i-- {
				if i+nums[now+i] > tempmax {
					tempmax = i + nums[now+i]
				} else {
					continue
				}
				if step+1 >= *minstep || nums[now+i] == 0 || i+nums[now+i] <= nums[now] {
					continue
				}
				findmin(nums, now+i, step+1, minstep)
			}
		}
	}

}

func findminDP(nums []int, now int, dp map[int]int) int {
	//fmt.Println(now)
	//fmt.Println(dp)
	if now >= len(nums)-1 {
		return 0
	}
	if dp[now] != 0 {
		return dp[now]
	}
	tempmin := 10000000
	for i := nums[now]; i > 0; i-- {
		if now+i < len(nums)-1 && (nums[now+i] == 0 || i+nums[now+i] <= nums[now]) {
			continue
		}
		tempstep := findminDP(nums, now+i, dp)
		if tempstep+1 < tempmin {
			tempmin = tempstep + 1
		}
	}
	dp[now] = tempmin
	return dp[now]

}

func jump(nums []int) int {
	//fmt.Println(nums)
	dp := make(map[int]int, len(nums))
	return findminDP(nums, 0, dp)

}
