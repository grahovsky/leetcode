package problems

/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	m := make(map[*ListNode]bool)
	for head.Next != nil {
		m[head] = true
		if m[head.Next] {
			return true
		}
		head = head.Next
	}
	return false
}

// @lc code=end
