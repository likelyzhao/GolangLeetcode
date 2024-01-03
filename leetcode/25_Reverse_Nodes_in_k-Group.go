package leecode

/*
	Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.

You may not alter the values in the list's nodes, only nodes themselves may be changed.

Example 1:

Input: head = [1,2,3,4,5], k = 2
Output: [2,1,4,3,5]
Example 2:

Input: head = [1,2,3,4,5], k = 3
Output: [3,2,1,4,5]

Constraints:

The number of nodes in the list is n.
1 <= k <= n <= 5000
0 <= Node.val <= 1000

Follow-up: Can you solve the problem in O(1) extra memory space?
*/
func swapNodebyN(ori *ListNode, n int) *ListNode {
	temp := ori
	var tempList []*ListNode
	for i := 0; i < n; i++ {
		temp = temp.Next
		tempList = append(tempList, temp)
	}
	//fmt.Println(ShowList(ori))
	head := temp.Next
	temp = ori
	for i := n - 1; i >= 0; i-- {
		temp.Next = tempList[i]
		temp = temp.Next
	}
	temp.Next = head
	return temp
}

func checkswapable(ori *ListNode, k int) bool {
	for i := 0; i <= k; i++ {
		if ori == nil {
			return false
		}
		ori = ori.Next
	}
	return true
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	//fmt.Println(ShowList(head))
	ori := head
	for i := 0; i < k; i++ {
		if ori == nil {
			return head
		}
		ori = ori.Next
	}
	// swap with head
	temp := head
	var tempList []*ListNode
	for i := 0; i < k; i++ {
		tempList = append(tempList, temp)
		temp = temp.Next
	}

	ptr := temp
	head = tempList[k-1]
	temp = head
	for i := k - 2; i >= 0; i-- {
		temp.Next = tempList[i]
		temp = temp.Next
	}
	temp.Next = ptr
	ptr = temp
	// swap with rest
	for {
		if !checkswapable(ptr, k) {
			break
		}
		ptr = swapNodebyN(ptr, k)
	}
	//swapNodebyN(head, k)
	//fmt.Println(ShowList(head))

	return head

}
