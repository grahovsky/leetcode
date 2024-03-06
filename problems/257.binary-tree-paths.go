package problems

import "strconv"

/*
 * @lc app=leetcode id=257 lang=golang
 *
 * [257] Binary Tree Paths
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
func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	if root != nil {
		val := strconv.Itoa(root.Val)
		child := append(binaryTreePaths(root.Left), binaryTreePaths(root.Right)...)
		if len(child) == 0 {
			res = append(res, val)
		} else {
			for _, s := range child {
				res = append(res, val+"->"+s)
			}
		}
	}

	return res
}

// @lc code=end
