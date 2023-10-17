/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

package problems

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// recursive
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list2.Next, list1)
	return list2
}

// fast
func MergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	merged := &ListNode{}
	result := merged

	for list1 != nil && list2 != nil {
		if list2.Val < list1.Val {
			merged.Next = list2
			list2 = list2.Next
		} else {
			merged.Next = list1
			list1 = list1.Next
		}
		merged = merged.Next
	}

	if list1 == nil {
		merged.Next = list2
	}
	if list2 == nil {
		merged.Next = list1
	}

	return result.Next
}

// @lc code=end
