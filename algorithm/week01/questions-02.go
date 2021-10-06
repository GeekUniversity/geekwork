package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Leetcode21: https://leetcode-cn.com/problems/merge-two-sorted-lists/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 设置head
	head := &ListNode{Val: 0, Next: nil}
	tmp := head
	// 两者都不空
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}
	// 其中一个不为空
	for l1 != nil {
		tmp.Next = l1
		l1 = l1.Next
		tmp = tmp.Next
	}
	for l2 != nil {
		tmp.Next = l2
		l2 = l2.Next
		tmp = tmp.Next
	}
	return head.Next
}

func main() {

	//l1 := &ListNode{Val: -9, Next: &ListNode{Val: 3, Next: nil}}
	//l2 := &ListNode{Val: 5, Next: &ListNode{Val: 7, Next: nil}}
	//
	//test := mergeTwoLists(l1, l2)
	//for test != nil {
	//	fmt.Println(test.Val)
	//	test = test.Next
	//}

	num := make([]int, 1, 3)
	num = append(num, 1)
	num = append(num, 2)
	num = append(num, 3)
	fmt.Println(num)
	fmt.Println(len(num))
	fmt.Println(cap(num))
	//num = append(num, 1)
	fmt.Println(num[:len(num)-1])
}
