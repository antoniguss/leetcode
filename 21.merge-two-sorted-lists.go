package main

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

	cur1, cur2 := list1, list2

	for cur1 != nil && cur2 != nil {

		if cur1 != nil && cur1.Val >= cur2.Val {
			cur.Val = cur1.Val
			cur.Next = &ListNode{}
			cur = cur.Next

			cur1 = cur1.Next
		} else {

			cur.Val = cur2.Val
			cur.Next = &ListNode{}
			cur = cur.Next

			cur1 = cur1.Next
		}

	}

	return head

}

// @leet end

func main() {

}

