package leecode

import "fmt"

func RevetConvert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}
	stdlen := numRows*2 - 2
	repeated := (len(s) - 1) / stdlen
	tail := len(s) - repeated*stdlen - 1
	endnum := repeated
	var overhalf int
	if tail >= numRows-1 {
		endnum++
		overhalf = 1
	}

	ends := s[len(s)-endnum : len(s)]
	endidx := len(s) - endnum

	r := make([]byte, len(s))

	r[numRows-1] = ends[0]
	for i := 1; i < len(ends); i++ {
		r[numRows+i*stdlen-1] = ends[i]
	}

	for i := 0; i < numRows-2; i++ {
		num := 2 * len(ends)
		if overhalf == 1 {
			if (tail - numRows - i) < 0 {
				num--
			}
		} else {
			if tail >= numRows-1-i-1 {
				num++
			}
		}
		layer := s[endidx-num : endidx]
		endidx = endidx - num

		r[numRows-i-1-1] = layer[0]
		if len(layer) >= 1 {
			r[numRows+i+1-1] = layer[1]
		}

		for j := 2; j < len(layer); j++ {
			fmt.Println(j)
			fmt.Println(numRows + (j/2)*stdlen + (j % 2 * 2) - i - 1)
			r[numRows+(j/2)*stdlen+((j%2*2)-1)*(i+1)-1] = layer[j]
		}
		fmt.Println(r)
	}

	layer := s[0:endidx]
	for i := 0; i < len(layer); i++ {
		r[i*stdlen] = layer[i]

	}

	// r[]

	return string(r)
	//	return "a"
}

func Convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}
	r := make([]byte, len(s))

	stdlen := numRows*2 - 2
	repeated := (len(s))/stdlen + 1
	stdlenup := 0
	idx := 0
	for i := 0; i < numRows; i++ {
		for step := 0; step < repeated+1; step++ {
			locnow := step*(stdlen+stdlenup) - i
			if stdlen != 0 {
				if locnow < len(s) && locnow >= 0 {
					r[idx] = s[locnow]
					idx++
				}
			}
			if stdlenup != 0 {
				locnow += stdlenup
				if locnow < len(s) && locnow > 0 {
					r[idx] = s[locnow]
					idx++
				}
			}
		}
		stdlen -= 2
		stdlenup += 2
		fmt.Println(r)

	}

	// r[]

	return string(r)
	//	return "a"
}
