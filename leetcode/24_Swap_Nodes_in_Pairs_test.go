package leecode

import "testing"

/*
Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)

Example 1:

Input: head = [1,2,3,4]
Output: [2,1,4,3]
Example 2:

Input: head = []
Output: []
Example 3:

Input: head = [1]
Output: [1]

Constraints:

The number of nodes in the list is in the range [0, 100].
0 <= Node.val <= 100
*/
func Test24(t *testing.T) {
	if ans := swapPairs(construclistfromarray([]int{1, 2, 3, 4})); !slicecompare(ShowList(ans),
		[]int{2, 1, 4, 3}) {
		t.Errorf("[]int{1, 2,3,4}  expected be []int{2, 1, 4, 3} but %v got", ShowList(ans))
	}
	if ans := swapPairs(construclistfromarray([]int{})); !slicecompare(ShowList(ans),
		[]int{}) {
		t.Errorf("[]int{}  expected be []int{} but %v got", ShowList(ans))
	}

	if ans := swapPairs(construclistfromarray([]int{1})); !slicecompare(ShowList(ans),
		[]int{1}) {
		t.Errorf("[]int{1}  expected be []int{1} but %v got", ShowList(ans))
	}
}
