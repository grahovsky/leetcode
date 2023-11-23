/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
 */

package problems

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricNodes(root.Left, root.Right)
}

func isSymmetricNodes(left *TreeNode, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}

	return left.Val == right.Val && isSymmetricNodes(left.Left, right.Right) && isSymmetricNodes(left.Right, right.Left)
}

// @lc code=end
