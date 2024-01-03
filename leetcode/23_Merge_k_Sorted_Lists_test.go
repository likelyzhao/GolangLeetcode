package leecode

import "testing"

/*
You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.



Example 1:

Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6
Example 2:

Input: lists = []
Output: []
Example 3:

Input: lists = [[]]
Output: []


Constraints:

k == lists.length
0 <= k <= 104
0 <= lists[i].length <= 500
-104 <= lists[i][j] <= 104
lists[i] is sorted in ascending order.
The sum of lists[i].length will not exceed 104.
*/

func Test23(t *testing.T) {

	if ans := mergeKLists([]*ListNode{construclistfromarray([]int{1, 4, 5}),
		construclistfromarray([]int{1, 3, 4}), construclistfromarray([]int{2, 6})}); !slicecompare(ShowList(ans),
		[]int{1, 1, 2, 3, 4, 4, 5, 6}) {
		t.Errorf("[]int{1, 4, 5} []int{1, 3, 4} []int{2,6} expected be [1, 1, 2, 3, 4, 4, 5, 6] but %v got", ShowList(ans))
	}
	if ans := mergeKLists([]*ListNode{}); !slicecompare(ShowList(ans),
		[]int{}) {
		t.Errorf(" expected be [] but %v got", ShowList(ans))
	}
	if ans := mergeKLists([]*ListNode{construclistfromarray([]int{})}); !slicecompare(ShowList(ans),
		[]int{}) {
		t.Errorf(" expected be [] but %v got", ShowList(ans))
	}

	if ans := mergeKLists([]*ListNode{construclistfromarray([]int{1, 2, 3}), construclistfromarray([]int{4, 5, 6, 7})}); !slicecompare(ShowList(ans),
		[]int{1, 2, 3, 4, 5, 6, 7}) {
		t.Errorf(" expected be [] but %v got", ShowList(ans))
	}
}
