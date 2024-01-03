package leecode

import (
	"testing"
)

/*
Given the head of a linked list, remove the nth node from the end of the list and return its head.

Example 1:


Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
Example 2:

Input: head = [1], n = 1
Output: []
Example 3:

Input: head = [1,2], n = 1
Output: [1]


Constraints:

The number of nodes in the list is sz.
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz


Follow up: Could you do this in one pass?
*/

func construclistfromarray(input []int) *ListNode {
	//var outputList ListNode
	if len(input) == 0 {
		return nil
	}

	ptrList := &ListNode{Val: -1000, Next: nil}
	res := ptrList
	//fmt.Println(input)
	for idx, item := range input {
		if idx == 0 {
			ptrList.Val = item
		} else {
			ptrList.Next = &ListNode{Val: item, Next: nil}
			ptrList = ptrList.Next

		}
	}
	//fmt.Println(ShowList(res))
	return res
}

func Test19(t *testing.T) {

	if ans := removeNthFromEnd(construclistfromarray([]int{1, 2, 3, 4, 5}), 2); !slicecompare(ShowList(ans),
		[]int{1, 2, 3, 5}) {
		t.Errorf("[]int{1, 2, 3, 4, 5}, 2 expected be 1, 2, 3, 5, but %v got", ShowList(ans))
	}

	if ans := removeNthFromEnd(construclistfromarray([]int{1}), 1); !slicecompare(ShowList(ans),
		[]int{}) {
		t.Errorf("[]int{1}, 1 expected be {} but %v got", ShowList(ans))
	}

	if ans := removeNthFromEnd(construclistfromarray([]int{1, 2}), 1); !slicecompare(ShowList(ans),
		[]int{1}) {
		t.Errorf("[1,2], 1 expected be [1], but %v got", ans)
	}
	if ans := removeNthFromEnd(construclistfromarray([]int{1, 2}), 2); !slicecompare(ShowList(ans),
		[]int{2}) {
		t.Errorf("[1,2], 2 expected be [2], but %v got", ans)
	}

}
