package leecode

import "fmt"

func getmedian(num []int, num2 int) float64 {

	if len(num) == 1 {
		return (float64(num[0]) + float64(num2)) / 2
	}

	if (len(num)+1)%2 == 0 {
		temp := (num[(len(num)-1)/2])
		tempafter := (num[(len(num)-1)/2+1])
		tempbefore := (num[(len(num)-1)/2-1])
		if temp > num2 && tempbefore > num2 {
			return (float64(temp) + float64(tempbefore)) / 2
		}

		if temp > num2 && tempbefore <= num2 {
			return (float64(temp) + float64(num2)) / 2
		}

		if temp <= num2 && tempafter > num2 {
			return (float64(num2) + float64(temp)) / 2
		}

		if temp <= num2 && tempafter <= num2 {
			return (float64(temp) + float64(tempafter)) / 2
		}
	} else {
		temp := (num[(len(num))/2-1])
		tempafter := (num[(len(num))/2])
		if temp >= num2 {
			return float64(temp)
		}
		if temp < num2 && tempafter >= num2 {
			return float64(num2)
		}

		if tempafter < num2 {
			return float64(tempafter)
		}
	}

	//	if len(num)%2 == 0 {
	//		return (float64(num[len(num)/2]) + float64(num[len(num)/2-1])) / 2
	//	}

	//	return float64(num[(len(num)-1)/2])
	return 0

}

func mediansigle(num []int) float64 {
	if len(num)%2 == 0 {
		return (float64(num[len(num)/2]) + float64(num[len(num)/2-1])) / 2
	}
	return float64(num[(len(num)-1)/2])
}

func getHighlowIdx(num []int) (high, low, idx int) {

	if len(num)%2 == 0 {
		idx = len(num) / 2
		high = num[idx]
		low = num[idx-1]
		return
	} else {
		idx = len(num) / 2
		high = num[idx]
		low = high
		return
	}
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var tempmedian float64
	//totallen := len(nums1) + len(nums2)
	fmt.Println(nums1)
	fmt.Println(nums2)
	if len(nums1) == 0 {
		return mediansigle(nums2)
	}

	if len(nums2) == 0 {
		return mediansigle(nums1)
	}

	if len(nums1) == 1 {
		return getmedian(nums2, nums1[0])
	}

	if len(nums2) == 1 {
		return getmedian(nums1, nums2[0])
	}

	if len(nums1) < len(nums2) {
		tempshorthigh, tempshortlow, idx := getHighlowIdx(nums1)
		templonglow := nums2[idx-1]
		templonghigh := nums2[len(nums2)-idx]
		if tempshorthigh >= templonghigh {
			nums1 = nums1[0:idx]
		} else {
			nums2 = nums2[0 : len(nums2)-idx]
		}

		if tempshortlow >= templonglow {
			nums2 = nums2[idx:len(nums2)]
		} else {
			nums1 = nums1[idx:len(nums1)]
			//		nums2 = nums2[idx:len(nums2)]
		}
		tempmedian = FindMedianSortedArrays(nums1, nums2)

	} else {
		tempshorthigh, tempshortlow, idx := getHighlowIdx(nums2)
		templonglow := nums1[idx-1]
		templonghigh := nums1[len(nums1)-idx]
		if tempshorthigh >= templonghigh {
			nums2 = nums2[0 : len(nums2)-idx]
		} else {
			nums1 = nums1[0 : len(nums1)-idx]
		}

		if tempshortlow >= templonglow {
			nums1 = nums1[idx:len(nums1)]
		} else {
			nums2 = nums2[idx:len(nums2)]
			//		nums2 = nums2[idx:len(nums2)]
		}
		tempmedian = FindMedianSortedArrays(nums1, nums2)
	}

	return tempmedian
}
