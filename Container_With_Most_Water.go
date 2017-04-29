package leecode

func MaxArea(height []int) int {

	var startidx, endidx int
	var maxarea int
	startidx = 0
	endidx = len(height) - 1
	for {
		if startidx > endidx {
			break
		}
		heightend := height[endidx]
		heightstart := height[startidx]
		height := heightstart
		if heightstart > heightend {
			height = heightend
		}
		if maxarea < height*(endidx-startidx) {
			maxarea = height * (endidx - startidx)
		}
		if heightstart > heightend {
			endidx--
		} else {
			startidx++
		}

		//maxarea = math.Max(math.MinInt32(heightstart, heightend)*(endidx-startidx), maxarea)
	}

	return maxarea
}
