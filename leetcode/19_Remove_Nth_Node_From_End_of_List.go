package leecode

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
func ShowList(head *ListNode) []int {
	var ans []int
	for head != nil {
		ans = append(ans, head.Val)
		head = head.Next
	}
	return ans
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil && n == 1 {
		return nil
	}
	ptrforward := head
	ptrnow := head
	for i := 0; i < n; i++ {
		ptrforward = ptrforward.Next
	}
	if ptrforward == nil {
		return head.Next
	}
	//fmt.Println(ptrforward)
	for ptrforward.Next != nil {
		ptrforward = ptrforward.Next
		ptrnow = ptrnow.Next
		//fmt.Println("forword", ptrforward)
		//fmt.Println(ptrnow)
	}
	if ptrnow.Next == nil {
		ptrnow.Next = nil
	} else {
		ptrnow.Next = ptrnow.Next.Next
	}

	//fmt.Println("head", head)
	return head
}
