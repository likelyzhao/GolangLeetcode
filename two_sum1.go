package leecode

var reschan chan int

func findnumidx(nums []int, target int, idxfirst int) int {
	res := -1
	for idx, v := range nums {
		if idxfirst == idx {
			continue
		}
		if v > target {
			return res
		}

		if v == target {
			reschan <- idx
			reschan <- idxfirst

			return res
		}
	}
	return res

}

func TwoSum(nums []int, target int) []int {
	//for()
	var res []int
	reschan = make(chan int, 2)
	//	sort.Ints(nums)
	for idx, v := range nums {
		go findnumidx(nums, target-v, idx)
	}

	res = append(res, <-reschan)
	res = append(res, <-reschan)
	return res
}
