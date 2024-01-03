package leecode

import "testing"

/*
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.



Example 1:


Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
Example 2:

Input: list1 = [], list2 = []
Output: []
Example 3:

Input: list1 = [], list2 = [0]
Output: [0]


Constraints:

The number of nodes in both lists is in the range [0, 50].
-100 <= Node.val <= 100
Both list1 and list2 are sorted in non-decreasing order.
*/

func Test21(t *testing.T) {

	if ans := mergeTwoLists(construclistfromarray([]int{1, 2, 4}), construclistfromarray([]int{1, 3, 4})); !slicecompare(ShowList(ans),
		[]int{1, 1, 2, 3, 4, 4}) {
		t.Errorf("[]int{1, 2, 4} []int{1, 3, 4} expected be [1, 1, 2, 3, 4, 4] but %v got", ShowList(ans))
	}

	if ans := mergeTwoLists(construclistfromarray([]int{}), construclistfromarray([]int{})); !slicecompare(ShowList(ans),
		[]int{}) {
		t.Errorf("[]int{} []int{} expected be [] but %v got", ShowList(ans))
	}
	if ans := mergeTwoLists(construclistfromarray([]int{}), construclistfromarray([]int{0})); !slicecompare(ShowList(ans),
		[]int{0}) {
		t.Errorf("[]int{} []int{0} expected be [0] but %v got", ShowList(ans))
	}
	if ans := mergeTwoLists(construclistfromarray([]int{2}), construclistfromarray([]int{1})); !slicecompare(ShowList(ans),
		[]int{1, 2}) {
		t.Errorf("[]int{2} []int{1} expected be [1,2] but %v got", ShowList(ans))
	}
	if ans := mergeTwoLists(construclistfromarray([]int{-9, 3}), construclistfromarray([]int{5, 7})); !slicecompare(ShowList(ans),
		[]int{-9, 3, 5, 7}) {
		t.Errorf("[]int{-9, 3} []int{5, 7} expected be [-9, 3, 5, 7] but %v got", ShowList(ans))
	}

}
