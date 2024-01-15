package leecode

import "testing"

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

func Test51(t *testing.T) {

	if ans := solveNQueens(4); !twoDstringslicecompare(ans, [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}) {
		t.Errorf(" 4  expected be {\".Q..\",\"...Q\",\"Q...\",\"..Q.\"}, but %v got", ans)
	}

	if ans := solveNQueens(1); !twoDstringslicecompare(ans, [][]string{{"Q"}}) {
		t.Errorf(" 1  expected be {\"Q\"}, but %v got", ans)
	}

}
