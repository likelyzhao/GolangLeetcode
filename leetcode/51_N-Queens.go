package leecode

import (
	"strings"
)

/*
The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.

Given an integer n, return all distinct solutions to the n-queens puzzle. You may return the answer in any order.

Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space, respectively.

Example 1:

Input: n = 4
Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above
Example 2:

Input: n = 1
Output: [["Q"]]

Constraints:

1 <= n <= 9
*/

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func checklocal(index, loc int, temphis []int, len int) bool {
	for i, v := range temphis {
		if v == loc || (abs(i-index) == abs(v-loc)) {
			return false
		}
	}
	return true
}

func checkqueen(indexNow int, temphis []int, len int, res *[][]string) {
	if indexNow > len {
		return
	}
	//fmt.Println(indexNow, len, temphis)
	for i := 0; i < len; i++ {
		//fmt.Println("indexNow=", indexNow, "i=", i, "temphis", temphis)
		if checklocal(indexNow, i, temphis[0:indexNow], len) {
			//temphis = append(temphis[0:indexNow], i)
			temphis[indexNow] = i
			//fmt.Println(temphis)
			if indexNow == len-1 {
				temp := make([]string, len)
				for i, v := range temphis {
					temp[i] = strings.Repeat(".", v) + "Q" + strings.Repeat(".", len-v-1)
				}
				//fmt.Println("addtemp ", temp)
				*res = append(*res, temp)
			} else {
				checkqueen(indexNow+1, temphis, len, res)
			}
		}
	}

}

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	checkqueen(0, make([]int, n), n, &res)
	//fmt.Println(len(res))
	return res

}
