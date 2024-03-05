package problems

/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	mid := getMid(head)
	tail := mid.Next
	mid.Next = nil

	tail = reverse(tail)
	return compare(head, tail)
}

func getMid(head *ListNode) *ListNode {
	last, mid := head, head
	for last.Next != nil && last.Next.Next != nil {
		last = last.Next.Next
		mid = mid.Next
	}
	return mid
}

func reverse(head *ListNode) *ListNode {
	var rev *ListNode

	for head != nil {
		next := head.Next
		head.Next = rev
		rev = head
		head = next
	}

	return rev
}

func compare(first, second *ListNode) bool {
	for first != nil && second != nil {
		if first.Val != second.Val {
			return false
		}
		first = first.Next
		second = second.Next
	}
	return true
}

// @lc code=end
