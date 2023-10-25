/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

package problems

// @lc code=start
func climbStairs(n int) int {
	cur, next := 0, 1
	for ; n > 0; n-- {
		cur, next = next, cur+next
	}
	return next
}

// @lc code=end
