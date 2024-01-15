package leecode

/*
You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

Example 1:

Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,4,1],[8,5,2],[9,6,3]]
Example 2:

Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

Constraints:

n == matrix.length == matrix[i].length
1 <= n <= 20
-1000 <= matrix[i][j] <= 1000
*/
type loc struct {
	x, y int
}

func getval(matrix [][]int, loc loc) int {
	return matrix[loc.x][loc.y]
}

func setval(matrix [][]int, loc loc, val int) {
	matrix[loc.x][loc.y] = val
}

func rotate(matrix [][]int) {
	//fmt.Println("input", matrix)
	size := len(matrix)
	length := size - 1
	for i := 0; i < size/2; i++ {
		// get
		//fmt.Println(matrix)
		startidx := []loc{{x: i, y: i}, {x: i + length, y: i}, {x: i + length, y: i + length}, {x: i, y: i + length}}
		//fmt.Println(length)

		for j := 0; j < length; j++ {
			//fmt.Println(startidx)
			//swap four location
			//if j == 1 {
			//	fmt.Println(getval(matrix, startidx[0]), getval(matrix, startidx[1]), getval(matrix, startidx[2]), getval(matrix, startidx[3]))
			//}
			temp := getval(matrix, startidx[0])
			setval(matrix, startidx[0], getval(matrix, startidx[1]))
			setval(matrix, startidx[1], getval(matrix, startidx[2]))
			setval(matrix, startidx[2], getval(matrix, startidx[3]))
			setval(matrix, startidx[3], temp)
			//update location
			startidx[0].y++
			startidx[1].x--
			startidx[2].y--
			startidx[3].x++
			//	fmt.Println("j=", j, matrix)
		}

		length -= 2
	}

}
