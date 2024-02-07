package problems

/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	diff := deep(root.Left) - deep(root.Right)
	return diff > -2 && diff < 2 && isBalanced(root.Left) && isBalanced(root.Right)
}

func deep(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(deep(root.Left), deep(root.Right)) + 1
}

// @lc code=end
