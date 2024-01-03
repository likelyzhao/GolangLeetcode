package leecode

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
func swapNode(ori, l1, l2 *ListNode) {
	temp := l2.Next
	ori.Next = l2
	l2.Next = l1
	l1.Next = temp
	return
}

func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	ptr := head
	temp := head.Next.Next
	head = head.Next
	head.Next = ptr
	ptr.Next = temp
	//fmt.Println(head.Val, head.Next.Val)
	//os.Exit(1)

	ptr = head.Next

	for ptr.Next != nil && ptr.Next.Next != nil {
		//fmt.Println(ptr.Val, ptr.Next.Val)
		swapNode(ptr, ptr.Next, ptr.Next.Next)
		ptr = ptr.Next.Next
		//fmt.Println(ptr.Val, ptr.Next.Val)
	}

	return head

}
