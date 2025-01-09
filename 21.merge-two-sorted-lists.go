package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// @leet start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	head := &ListNode{}
	cur := head

	for list1 != nil && list2 != nil {

		if list1.Val <= list2.Val {
			cur.Next = &ListNode{list1.Val, nil}
			cur = cur.Next

			list1 = list1.Next

		} else {
			cur.Next = &ListNode{list2.Val, nil}
			cur = cur.Next

			list2 = list2.Next
		}

	}

	for list1 != nil {
		cur.Next = &ListNode{list1.Val, nil}
		cur = cur.Next

		list1 = list1.Next

	}

	for list2 != nil {
		cur.Next = &ListNode{list2.Val, nil}
		cur = cur.Next

		list2 = list2.Next

	}

	return head.Next

}

// @leet end

func main() {

}
