package leecode

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
func checknode(lists []*ListNode) bool {
	for _, v := range lists {
		if v != nil {
			return true
		}
	}
	return false
}

func insertbyValue(lists []*ListNode, node *ListNode) []*ListNode {
	if node == nil {
		lists = append(lists, node)
		return lists
	}
	for idx, v := range lists {
		if v != nil && v.Val > node.Val {
			lists = append(lists[:idx], append([]*ListNode{node}, lists[idx:]...)...)
			return lists
		}
		if v == nil {
			lists = append(lists[:idx], append([]*ListNode{node}, lists[idx:]...)...)
			return lists
		}
		//fmt.Println(v)
	}
	lists = append(lists, node)
	return lists

}

func showlistarray(lists []*ListNode) []int {
	var ans []int
	for _, val := range lists {
		if val != nil {
			ans = append(ans, val.Val)
		} else {
			ans = append(ans, 10000)
		}

	}
	return ans
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var cadidateNode []*ListNode
	var head, ptr *ListNode
	ptr = nil
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			cadidateNode = insertbyValue(cadidateNode, lists[i])
		}
	}
	//fmt.Println(showlistarray(cadidateNode))
	for checknode(cadidateNode) {
		if ptr == nil {
			ptr = cadidateNode[0]
			head = ptr
		} else {
			ptr.Next = cadidateNode[0]
			ptr = ptr.Next
			//cadidateNode = insertbyValue(cadidateNode[1:], cadidateNode[0].Next)
		}

		insertnode := cadidateNode[0].Next
		cadidateNode = insertbyValue(cadidateNode[1:], insertnode)
		//fmt.Println("output", showlistarray(cadidateNode))

	}

	return head

}

func mergeKListsback(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	head := lists[0]
	for i := 0; i < len(lists)-1; i++ {
		head = mergeTwoLists(head, lists[i+1])
	}

	return head

}
