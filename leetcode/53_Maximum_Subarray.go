package leecode

/*
Given an integer array nums, find the
subarray
 with the largest sum, and return its sum.



Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.
Example 2:

Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.
Example 3:

Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.


Constraints:

1 <= nums.length <= 105
-104 <= nums[i] <= 104


Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
*/

func maxSubArray(nums []int) int {
	//fmt.Println("nums", nums)
	var lastsum, discardsum, newsum int
	lastsum = nums[0]
	discardsum = 0
	newsum = 0
	totalmaxsum := lastsum
	for idx, v := range nums {
		if idx == 0 {
			continue
		}
		//fmt.Println(lastsum, discardsum, newsum, v, totalmaxsum)
		if lastsum < 0 {
			if v > lastsum {
				lastsum = v
			}
			continue
		} else {
			if newsum > 0 {
				if v < 0 {
					if newsum+v > 0 {
						newsum += v
					} else {
						discardsum = discardsum + newsum + v
						newsum = 0
					}
				} else {
					newsum += v
				}
			} else if discardsum == 0 {
				if v < 0 {
					discardsum += v
				} else {
					lastsum += v
				}
			} else {
				if v < 0 {
					discardsum += v
				} else {
					newsum += v
				}
			}

			//fmt.Println(lastsum, discardsum, newsum, v, totalmaxsum)
			if lastsum >= 0 && discardsum < 0 && newsum >= 0 {
				if lastsum+discardsum < 0 && newsum+discardsum < 0 {
					if lastsum > totalmaxsum {
						totalmaxsum = lastsum
					}
					lastsum = newsum
					discardsum = 0
					newsum = 0

				}

				if newsum+discardsum > 0 {
					if newsum > lastsum+newsum+discardsum {
						lastsum = newsum
					} else {
						lastsum = lastsum + newsum + discardsum
					}
					discardsum = 0
					newsum = 0
				}
			}
		}
	}
	if lastsum > totalmaxsum {
		return lastsum
	}

	return totalmaxsum
}
