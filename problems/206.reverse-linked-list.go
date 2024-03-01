package problems

/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var res *ListNode

	for head != nil {
		res = &ListNode{Val: head.Val, Next: res}
		head = head.Next
	}

	return res
}

// @lc code=end
