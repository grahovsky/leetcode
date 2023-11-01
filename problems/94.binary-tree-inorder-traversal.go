/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
 */

package problems

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	nums := []int{}

	if root != nil {
		nums = append(nums, inorderTraversal(root.Left)...)
		nums = append(nums, root.Val)
		nums = append(nums, inorderTraversal(root.Right)...)
	}

	return nums
}

// @lc code=end
