package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// @leet start

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var c, v1, v2, s int
	dummyNode := &ListNode{}
	n := dummyNode

	for l1 != nil || l2 != nil || c > 0 {
		v1, v2 = 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		s = v1 + v2 + c
		c = s / 10

		newNode := &ListNode{Val: s % 10}
		n.Next = newNode
		n = n.Next

	}

	return dummyNode.Next
}

// @leet end
