/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package leecode

import (
	"container/list"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var addflag int
	var add1, add2 int
	addflag = 0
	listres := list.New()
	listp1 := l1
	listp2 := l2
	add1 = listp1.Val
	add2 = listp2.Val
	for {
		tempres := 0
		tempres += add1
		tempres += add2
		tempres += addflag
		if tempres >= 10 {
			addflag = 1
			listres.PushBack(tempres - 10)
		} else {
			addflag = 0
			listres.PushBack(tempres)
		}
		if listp2.Next == nil && listp1.Next == nil {
			break
		}
		if listp1.Next == nil {
			add1 = 0
			continue
		} else {
			listp1 = listp1.Next
			add1 = listp1.Val
		}
		if listp2.Next == nil {
			add2 = 0
			continue
		} else {
			listp2 = listp2.Next
			add2 = listp2.Val
		}

	}

	if addflag == 1 {
		listres.PushBack(1)
	}

	var res []ListNode
	valend := listres.Front()
	for {
		//	var tempres ListNode

		//	ListNode.new()
		tempres := ListNode{}
		tempres.Val = valend.Value.(int)
		//	tempres.Next = 00
		if len(res) == 0 {
			res = append(res, tempres)
		} else {
			res = append(res, tempres)
			//	res[len(res)-2].Next = &res[len(res)-1]
		}
		valend = valend.Next()
		if valend == nil {
			break
		}

	}

	for i := 0; i < len(res)-1; i++ {
		res[i].Next = &res[i+1]
	}

	var listpp *ListNode
	listpp = &res[0]
	for {
		fmt.Println(listpp.Val)
		if listpp.Next == nil {
			break
		}
		listpp = listpp.Next
	}
	return &res[0]
}
